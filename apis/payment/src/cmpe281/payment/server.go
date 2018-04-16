package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func main() {
	var wait time.Duration
	var dbuser, dbpass, dbkeyspace, dbhosts string
	var debug bool
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dbuser, "dbuser", "starbucks",
		"the username for cassandra database connection")
	flag.StringVar(&dbpass, "dbpass", "",
		"the password for cassandra database connection")
	flag.StringVar(&dbkeyspace, "dbkeyspace", "starbucks",
		"the keyspace for cassandra database connection")
	flag.StringVar(&dbhosts, "dbhosts", "54.176.100.87,54.241.192.98",
		"the hosts (comma separated) for cassandra database connection")
	flag.BoolVar(&debug, "debug", false, "run server in debug mode")
	flag.Parse()

	dbhostsparsed := strings.Split(dbhosts, ",")
	for i := 0; i < len(dbhostsparsed); i++ {
		dbhostsparsed[i] = strings.TrimSpace(dbhostsparsed[i])
	}

	cluster := gocql.NewCluster(parseDatabaseHosts(dbhosts)...)
	cluster.Keyspace = dbkeyspace
	cluster.Timeout = 5 * time.Second
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbuser,
		Password: dbpass,
	}

	log.Printf("Connecting to Cassandra...")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("Done")

	server := &Server{
		debug:     debug,
		cassandra: session,
	}

	router := mux.NewRouter()
	router.Use(server.AuthMiddleware)
	router.HandleFunc("/payments", server.ListPayments)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until SIGINT (Ctrl+C) is received, then begin shutdown.
	<-c

	log.Printf("Server beginning shutdown")
	session.Close()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Wait until all connections are finished or until timeout
	srv.Shutdown(ctx)
	log.Println("Server shut down successfully")
	os.Exit(0)
}

// Authenticate the User, add 'user_id' to request context
func (srv *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authentication")

		user_id, found := getUser(token)
		if !found && srv.debug {
			// Allow passing of User ID directly if in debug mode
			user_id = r.Header.Get("X-User-ID")
		}

		if user_id != "" {
			// If User ID was found, continue as Authenticated User
			log.Printf("Authenticated User [%s]", user_id)
			user_context := context.WithValue(r.Context(), "user_id", user_id)
			next.ServeHTTP(w, r.WithContext(user_context))
		} else {
			// Else throw error for Unauthenticated User
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// Retrieve User ID from Authentication Token
func getUser(token string) (string, bool) {
	// TODO(bbamsch): Implement Token->UserID Logic
	// return "a77be83b-f9c1-4087-a81b-511f3ad2c9cb", true
	return "", false
}

// Takes a string of comma separated database hosts and
// returns them as an array of database hosts
func parseDatabaseHosts(hosts string) []string {
	parsedhosts := strings.Split(hosts, ",")
	for i := 0; i < len(parsedhosts); i++ {
		parsedhosts[i] = strings.TrimSpace(parsedhosts[i])
	}
	return parsedhosts
}

type Server struct {
	debug     bool
	cassandra *gocql.Session
}

type PaymentDetails struct {
	PaymentId gocql.UUID `json:"payment_id"`
	Timestamp time.Time  `json:"timestamp"`
	Amount    float64    `json:"amount"`
}

type ListPaymentsResult struct {
	Payments      []*PaymentDetails `json:"payments"`
	NextPageToken *gocql.UUID       `json:"next_page_token"`
}

func (srv *Server) ListPayments(w http.ResponseWriter, r *http.Request) {
	// Parse Query Parameters
	var limit uint64
	var pageToken string
	{
		vars := r.URL.Query()
		if query_limit, err := strconv.ParseUint(vars.Get("limit"), 10, 64); err == nil {
			limit = query_limit
		} else {
			limit = 10 // Default Limit
		}
		pageToken = vars.Get("pageToken")
	}

	// Run Query against DB
	user_id := r.Context().Value("user_id")
	query := "SELECT payment_id, amount FROM payments WHERE user_id = ? AND payment_id > ? LIMIT ?"
	payment_iter := srv.cassandra.Query(query).Bind(user_id, pageToken, limit).PageSize(10).Iter()

	// Extract Payment Details from Query Iterator
	var payments = make([]*PaymentDetails, 0)
	var paymentId *gocql.UUID = nil
	var amount *float64 = nil
	for payment_iter.Scan(paymentId, amount) {
		payments = append(payments, &PaymentDetails{
			PaymentId: *paymentId,
			Timestamp: paymentId.Time(),
			Amount:    *amount,
		})
	}

	// Build Response Structure
	result := &ListPaymentsResult{
		Payments:      payments,
		NextPageToken: paymentId,
	}

	// Transform Output to JSON
	if output, err := json.Marshal(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}

}