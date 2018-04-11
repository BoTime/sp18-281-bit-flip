package main

import (
	"fmt"
	"context"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
	"time"
	"log"
	"os"
	"strings"
	"os/signal"
)

func main() {
	var wait time.Duration
	var dbuser, dbpass, dbkeyspace, dbhosts string
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dbuser, "dbuser", "starbucks",
		"the username for cassandra database connection")
	flag.StringVar(&dbpass, "dbpass", "",
		"the password for cassandra database connection")
	flag.StringVar(&dbkeyspace, "dbkeyspace", "starbucks",
		"the keyspace for cassandra database connection")
	flag.StringVar(&dbhosts, "dbhosts", "54.176.100.87,54.241.192.98",
		"the hosts (comma separated) for cassandra database connection")
	flag.Parse()

	dbhostsparsed := strings.Split(dbhosts, ",")
	for i := 0; i < len(dbhostsparsed); i++ {
		dbhostsparsed[i] = strings.TrimSpace(dbhostsparsed[i])
	}

	cluster := gocql.NewCluster(ParseDatabaseHosts(dbhosts)...)
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

	var id gocql.UUID
	var firstname, lastname, gender string
	var query = "SELECT id, firstname, lastname, gender FROM test";
	if err = session.Query(query).Scan(&id, &firstname, &lastname, &gender); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("ID: %s, FirstName: %s, LastName: %s, Gender: %s", id, firstname, lastname, gender)

	router := mux.NewRouter()
	router.HandleFunc("/payment", Payment)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Second * 60,
		Handler: router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func ParseDatabaseHosts(hosts string) []string {
	parsedhosts := strings.Split(hosts, ",")
	for i := 0; i < len(parsedhosts); i++ {
		parsedhosts[i] = strings.TrimSpace(parsedhosts[i])
	}
	return parsedhosts
}

func Payment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "asdfjkl: %s", vars["test"])
}