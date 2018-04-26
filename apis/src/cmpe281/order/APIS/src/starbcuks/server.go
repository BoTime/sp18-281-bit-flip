/*
	Starbcuks order backend orders api
*/

package main 

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"time"
	"log"
	"strings"
)

var cluster *gocql.ClusterConfig

var hmacSecret = "bit-flip"
//Functions to decode AUht token to fetch user id
func getUser(rawToken string) (string, bool) {
	fields := strings.Fields(rawToken)
	if len(fields) < 1 {
		log.Println("User Auth Token Missing")
		return "", false
	}

	switch fields[0] {
	case "jwt":
		return getUserJwt(fields)
	default:
		log.Println("Unexpected User Auth Token Type")
		return "", false
	}
}

func getUserJwt(fields []string) (string, bool) {
	if len(fields) != 2 {
		log.Println("Invalid JWT Token")
		return "", false
	}

	token, err := jwt.Parse(fields[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("bit-flip"), nil
	})

	if err != nil {
		log.Println("User Auth Token failed Parsing")
		return "", false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"]
		if userId != nil {
			return userId.(string), true
		}
		log.Println("User ID not in Claims")
		return "", false
	} else {
		log.Println("User Auth Token failed Validation")
		return "", false
	}
}

func RemoveIndex(s []OrderDetails, index int) []OrderDetails {
    return append(s[:index], s[index+1:]...)
}
func main() {

	// connect to the cluster

	cluster = gocql.NewCluster("13.57.254.48")//,"54.183.23.103","54.183.136.246","54.67.64.236","54.193.97.78")
	cluster.Keyspace ="starbucks"
	cluster.ProtoVersion = 3
	cluster.DisableInitialHostLookup  = true
	cluster.Port = 9042
	cluster.Timeout = 10 * time.Second
	var port = "3000"
	go func(){
		for{
			fmt.Println("Sleeping....")
			time.Sleep(5 * time.Second)
			fmt.Println("Awake....")
			session, errs := cluster.CreateSession()
			if errs != nil{
				fmt.Println("ERROR session creation", errs)
			} else{
				fmt.Println("session created")
				defer session.Close()
				var userid, payid gocql.UUID
				var count int = 0
				if err := session.Query(`select count(*), user_id, pay_id from starbucks.orders where status ='placed' limit 1 ALLOW FILTERING` ).Scan(&count,&userid, &payid); err != nil {
					fmt.Println("ERROR in select first row", err)
				}else{
					if count > 0{
						fmt.Println("GOT",userid ,payid )
						fmt.Println("Processing ..")
						
						//Update status
						
						if err := session.Query(`update orders set status ='processed' where user_id = ? and pay_id = ?`,userid,payid ).Exec(); err != nil {
							fmt.Println("ERROR in update in Processing", err)
						}else{
							fmt.Println("Sucessfully Processed")
						}
					}	
				}
			}
		}	
	}()
	server := NewServer()
	server.Run(":" + port)
}

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", pingLBHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", pingAuthHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", starbcuksNewOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/orders", starbcuksOrderStatusHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", starbcuksDeleteOrdersHandler(formatter)).Methods("DELETE")
	//mx.HandleFunc("/orders", starbcuksProcessOrdersHandler(formatter)).Methods("POST")
}

// API Ping Handler LB
func pingLBHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("INSIDE PING" )
		var msg ="API ORDER ALIVE!"
		formatter.JSON(w, http.StatusOK, msg)
	}
}

// API Ping for Auth
func pingAuthHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//token := req.Header.Get("Authorization")
		token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		log.Println(token)
		fmt.Println("INSIDE PINGAuth" )
		_, found := getUser(token)
		if found{
			formatter.JSON(w, http.StatusOK, "Authenticated")
		}else{
			formatter.JSON(w, http.StatusBadRequest, "Bad Authentication")
		}
	}
}

// API Create New Starbcuks Order //TODO
func starbcuksNewOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Read body
		log.Println(req.Body)
		//token := req.Header.Get("Authorization")
		//jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI'
		token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		log.Println(token)
		log.Println("REQQBODY",req)
		userId, found := getUser(token)
		if found{
			var payload Order
			decoder := json.NewDecoder(req.Body)
			if err := decoder.Decode(&payload); err != nil {
				log.Println(err)
				formatter.JSON(w, http.StatusBadRequest, "Something went bad with payload")
				return
			}
			log.Println("BEFORE EMPTY")
			log.Println(payload)
			for i := 0; i < len(payload.Product); i++ {
				if payload.Product[i]== (OrderDetails{}){
					payload.Product = RemoveIndex(payload.Product,i)
					i= i-1
				}
					
			}
			log.Println("AFTER EMPTY")
			log.Println(payload)
			
			//id timeuuid, uuid,status text,pay_id uuid,store text, product frozen<items>,payment frozen<payments>,PRIMARY KEY (user_id, id)
			payload.UserId, _ = gocql.ParseUUID(userId)
			payload.PayId = gocql.TimeUUID()
			payload.Status = "placed"
			log.Println(payload)
			session, errs := cluster.CreateSession()
			if errs != nil{
				msg := "Error creating session"
				fmt.Println("ERROR session creation", errs)
				formatter.JSON(w, http.StatusInternalServerError, msg)
			} else{
				query, names := qb.Insert("orders").Columns("product", "payment", "user_id", "pay_id", "status", "store").ToCql()
				q := gocqlx.Query(session.Query(query), names).BindStruct(payload)
				
				if err := q.ExecRelease(); err != nil {
					log.Println("ERROR SAVING DATA")
					log.Println(err)
					formatter.JSON(w, http.StatusInternalServerError, "Failed to create Payment")
					return
				}
				formatter.JSON(w, http.StatusOK, payload )
			}
		}else{
			formatter.JSON(w, http.StatusBadRequest, "Bad Authentication")
		}
	}	
}

// API Get ALL Orders for a userID
func starbcuksOrderStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//token := req.Header.Get("Authorization")
		token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		
		userIdS, found := getUser(token)
		if found{			
			fmt.Println("GET ORDERS")
			fmt.Println("Userid", userIdS)
			var order []GetOrder
			session, errs := cluster.CreateSession()
			if errs != nil{
				msg := "Error creating session"
				fmt.Println("ERROR session creation", errs)
				formatter.JSON(w, http.StatusInternalServerError, msg)
				return
			} else{
				if user_id, err := gocql.ParseUUID(userIdS); err == nil {
					queryMap := qb.M{"user_id":nil}
					queryMap["user_id"] = user_id
					defer session.Close()
					stmt, names := qb.Select("orders").Where(qb.Eq("user_id")).ToCql()
					fmt.Println("querymap", queryMap)
					q := gocqlx.Query(session.Query(stmt), names).BindMap(queryMap)
					if err := gocqlx.Select(&order, q.Query); err != nil {
						fmt.Println("ERROR in select query execution", err)
						formatter.JSON(w, http.StatusInternalServerError, err)
						return
					}
					data, _:= json.Marshal(order)
					formatter.JSON(w, http.StatusOK, string(data))
				} else {
					formatter.JSON(w, http.StatusUnauthorized, "Unable to Authenticate")
				}
				
			}
		}else{
			formatter.JSON(w, http.StatusBadRequest, "Bad Authentication")
		}
	}	
}

// API Process Orders
func starbcuksProcessOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		/*for key, value := range orders {
			fmt.Println("Key:", key, "Value:", value)
			var ord = orders[key]
			ord.OrderStatus = "Order Processed"
			orders[key] = ord
		}
		fmt.Println("Orders: ", orders)*/
		formatter.JSON(w, http.StatusOK, "Orders Processed!")
	}
}
// API Delete an Order
func starbcuksDeleteOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//token := req.Header.Get("Authorization")
		token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		
		userIdS, found := getUser(token)
		if found{			
			fmt.Println("DELETE ORDERS")
			fmt.Println("Userid", userIdS)
			var pay Deletepay
			decoder := json.NewDecoder(req.Body)
			if err := decoder.Decode(&pay); err != nil {
				log.Println(err)
				formatter.JSON(w, http.StatusBadRequest, "Something went bad with payload")
				return
			}
			log.Println(pay.Pid)
			pay_idS := pay.Pid 
			fmt.Println(pay_idS)
			
			session, errs := cluster.CreateSession()
			if errs != nil{
				msg := "Error creating session"
				fmt.Println("ERROR session creation", errs)
				formatter.JSON(w, http.StatusInternalServerError, msg)
				return
			} else{
				if user_id, err := gocql.ParseUUID(userIdS); err == nil {
					if pay_id, err := gocql.ParseUUID(pay_idS); err == nil {								
						queryMap := qb.M{"pay_id": nil,"user_id":nil}
						queryMap["pay_id"] = pay_id
						queryMap["user_id"] = user_id
						var count = 0
						defer session.Close()
						if err := session.Query(`select count(*) from orders where user_id = ? and pay_id = ?`,user_id,pay_id ).Scan(&count); err != nil {
							fmt.Println("ERROR in select count(*) query execution", err)
							formatter.JSON(w, http.StatusInternalServerError, err)
							return
						}
						if count == 0{
							fmt.Println("QUery returned 0 count")
							formatter.JSON(w, http.StatusBadRequest, "Bad Payload")
							return
						}
						if err := session.Query(`delete from orders where user_id = ? and pay_id = ?`,user_id,pay_id ).Exec(); err != nil {
							fmt.Println("ERROR in delete query execution", err)
							formatter.JSON(w, http.StatusInternalServerError, err)
							return
						}
						fmt.Println("Sucessfully deleted")
						formatter.JSON(w, http.StatusAccepted, "Deleted")
					}else{
						fmt.Println("ERROR in parsing Paydetail", err)
						formatter.JSON(w, http.StatusBadRequest, err)
						return
					}
				}else{
					fmt.Println("ERROR in parsing User detail", err)
					formatter.JSON(w, http.StatusUnauthorized, "Unable to Authenticate")
					return
				}
			}
		}else{
			formatter.JSON(w, http.StatusBadRequest, "Bad Authentication")
		}
	}
}	