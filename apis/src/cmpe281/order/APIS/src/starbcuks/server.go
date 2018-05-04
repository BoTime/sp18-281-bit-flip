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
	"bytes"
	"io/ioutil"
)

var cluster *gocql.ClusterConfig
var pay_url = "http://kong-lb-133222058.us-west-1.elb.amazonaws.com/payments/v1/payments"
var invtry_url = "http://kong-lb-133222058.us-west-1.elb.amazonaws.com/inventory/v1/stores/"
var hmacSecret = "bit-flip"

//Functions to decode Auth token to fetch user id
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
//Function to call PAY API to created PAYID
func PostPay(token string, payload Payments)(int, GetPayments){
    log.Println("Post Pay function BEGIN")
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)
	log.Println("Pay Rqst BODY", b)
	//To hold payment response
	var payresp GetPayments
    req, err := http.NewRequest("POST", pay_url, b)
    req.Header.Set("Authorization",token)
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
	log.Println("PAYEMNT API CALLED, reqst: ", req)
    if err != nil {
		log.Println("ERROR IN POST PAY", err)
        return 404, payresp
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &payresp)
	log.Println("UNMARSHALLED RESPONSE: ", payresp)
	log.Println("Inside Post Pay function END")
	return resp.StatusCode, payresp
}

//Function to call Inventory API to check inventory stock
func PostInventory(token string, store string, payload Order)(int, string, gocql.UUID ){
    log.Println("Inside Post Inventory function BEGIN")
	
	//Struct to create json for request
	type TempProducts struct {
		Products []OrderDetails `json:"products"`
	}
	var prod = TempProducts{Products : payload.Product}
	log.Println("Request body struct:", prod)
	
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(prod)
    req, err := http.NewRequest("POST", invtry_url+store+"/allocations", b)
    req.Header.Set("Authorization",token)
    req.Header.Set("Content-Type", "application/json")
	//To store response of API
	var invresp GetInventory
    client := &http.Client{}
    resp, err := client.Do(req)
	log.Println("Inventory API CALLED, reqst: ", req)
	
    if err != nil {
		log.Println("ERROR IN Inventory allocation", err)
        return 404, invresp.Status, invresp.InvId
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &invresp)
	log.Println("UNMARSHALLED RESPONSE: ", resp)
	log.Println("Inside Post Inventory function END")
	return resp.StatusCode, invresp.Status, invresp.InvId
}

//Function to call Inventory confirmation API
func ConfirmInventory(token string, store string, allocId string )(int, string){
    log.Println("Inside COnfirmation Inventory function END")
    req, err := http.NewRequest("POST", invtry_url+store+"/allocations/"+allocId, nil)
    req.Header.Set("Authorization",token)
    req.Header.Set("Content-Type", "application/json")
	var invresp GetInventory
    client := &http.Client{}
    resp, err := client.Do(req)
	log.Println("Inventory confirmation API CALLED", req)
    if err != nil {
		log.Println("ERROR IN Inventory confirmation", err)
        return 404, invresp.Status
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &invresp)
	log.Println("UNMARSHALLED RESPONSE", invresp)
    log.Println("Inside COnfirmation Inventory function BEGIN")
	return resp.StatusCode, invresp.Status
}


func RemoveIndex(s []OrderDetails, index int) []OrderDetails {
    return append(s[:index], s[index+1:]...)
}
func main() {

	// connect to the cluster

	cluster = gocql.NewCluster("13.57.20.19","13.56.255.172","13.56.207.85","54.215.219.48","54.193.109.49")
	cluster.Keyspace ="starbucks"
	cluster.ProtoVersion = 3
	cluster.DisableInitialHostLookup  = true
	cluster.Port = 9042
	cluster.Timeout = 10 * time.Second
	var port = "3000"
	
	// Routine to pick one placed orders from database every 1 minute to update it as processed
	go func(){
		for{
			log.Println("Sleeping....")
			time.Sleep(120 * time.Second)
			log.Println("Awake....")
			session, errs := cluster.CreateSession()
			if errs != nil{
				log.Println("ERROR session creation", errs)
			} else{
				log.Println("session created")
				defer session.Close()
				var userid, payid gocql.UUID
				var count int = 0
				if err := session.Query(`select count(*), user_id, pay_id from starbucks.orders where status ='placed' limit 1 ALLOW FILTERING` ).Scan(&count,&userid, &payid); err != nil {
					log.Println("ERROR in select first row", err)
				}else{
					if count > 0{
						log.Println("GOT",userid ,payid )
						log.Println("Processing ..")			
						
						//Update status						
						if err := session.Query(`update orders set status ='processed' where user_id = ? and pay_id = ?`,userid,payid ).Exec(); err != nil {
							log.Println("ERROR in update in Processing", err)
						}else{
							log.Println("Sucessfully Processed")
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
	mx.HandleFunc("/ping", pingAuthHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", starbcuksNewOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/order", starbcuksOrderStatusHandler(formatter)).Methods("GET")
	mx.HandleFunc("/order", starbcuksDeleteOrdersHandler(formatter)).Methods("DELETE")
}

// API Ping Handler LB
func pingLBHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("INSIDE PING" )
		var msg ="API ORDER ALIVE!"
		formatter.JSON(w, http.StatusOK, msg)
	}
}

// API Ping for Auth
func pingAuthHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("INSIDE PING Auth" )
		token := req.Header.Get("Authorization")
		//token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		log.Println(token)
		_, found := getUser(token)
		if found{
			formatter.JSON(w, http.StatusOK, "Authenticated")
		}else{
			formatter.JSON(w, http.StatusUnauthorized, "Bad Authentication")
		}
	}
}

// API Create New Starbcuks Order 
func starbcuksNewOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Read body
		log.Println("Inside Order create BEGIN")
		log.Println("REQUEST BODY:", req.Body)
		token := req.Header.Get("Authorization")
		//token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		log.Println(token)
		userId, found := getUser(token)
		if found{
			var payload Order //To store request
			decoder := json.NewDecoder(req.Body)
			if err := decoder.Decode(&payload); err != nil {
				log.Println(err)
				formatter.JSON(w, http.StatusInternalServerError, "Something went bad with payload")
				return
			}
			// To remove any empty Product struct in request
			log.Println("Removing any EMPTY products from request...")
			log.Println("Before:",payload)
			for i := 0; i < len(payload.Product); i++ {
				if payload.Product[i]== (OrderDetails{}){
					payload.Product = RemoveIndex(payload.Product,i)
					i= i-1
				}					
			}
			log.Println("After:", payload)
			
			//Call Inventory
			log.Println("Calling Inventory POST")			
			code, allstatus, invid := PostInventory(token,payload.Store, payload)
			log.Println("Inventory POST Fucntion resp: ", code)
			if code != 404 && code != 400 && code != 500 && code != 503 {
				if allstatus == "Unconfirmed"{
					log.Println("Sucess calling Inventory")
					//Call Payment
					log.Println("Calling Payment POST")
					paystatus, pay := PostPay(token, payload.Payment)
					if paystatus != 404 && paystatus != 400  && paystatus != 503 && paystatus != 500{
						if pay.Status == "Declined"{
							log.Println("Payment status Declined")
							formatter.JSON(w, http.StatusInternalServerError, "Bad Card Details")
							return
						}
						log.Println("Success Pay API")
						payload.UserId, _ = gocql.ParseUUID(userId)
						payload.PayId = pay.PayID
						payload.Status = "placed"
						log.Println("Creating order:",payload)
						session, errs := cluster.CreateSession()
						if errs != nil{
							msg := "Error creating session"
							log.Println("ERROR session creation", errs)
							formatter.JSON(w, http.StatusInternalServerError, msg)
							return
						} else{
							query, names := qb.Insert("orders").Columns("product", "payment", "user_id", "pay_id", "status", "store").ToCql()
							q := gocqlx.Query(session.Query(query), names).BindStruct(payload)
							
							if err := q.ExecRelease(); err != nil {
								log.Println("ERROR SAVING ORDER DATA", err)
								formatter.JSON(w, http.StatusInternalServerError, "Failed to create order")
								return
							}
							log.Println("Sucess creating order, confirming Inventory allocation:",invid)
							
							//Call Allocation Confirmation-TODO
							log.Println("CAlled Allocation COnformation:", invid)
							_, allstatus := ConfirmInventory(token,payload.Store, invid.String())
							if allstatus == "Confirmed" {
								formatter.JSON(w, http.StatusCreated, "Created")
								return
							}else{
								log.Println("Allocation COnformation: threw status ", allstatus)
							}
						}
					}else{
						formatter.JSON(w, http.StatusInternalServerError, "Bad Dependency-Payment")
					}
				}else{
					formatter.JSON(w, http.StatusInternalServerError, "Inventory Lacking")
					return
				}
			}else{
				formatter.JSON(w, http.StatusInternalServerError, "Bad Dependency-Inventory")
				return
			}			
		}else{
			formatter.JSON(w, http.StatusUnauthorized, "Bad Authentication")
		}
	}	
}

// API Get ALL Orders for a userID
func starbcuksOrderStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		//token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
		
		userIdS, found := getUser(token)
		if found{			
			log.Println("GET ORDERS BEGIN for ID:",userIdS)
			var order []GetOrder
			session, errs := cluster.CreateSession()
			if errs != nil{
				msg := "Error creating session"
				log.Println("ERROR session creation", errs)
				formatter.JSON(w, http.StatusInternalServerError, msg)
				return
			}else{
				if user_id, err := gocql.ParseUUID(userIdS); err == nil {
					queryMap := qb.M{"user_id":nil}
					queryMap["user_id"] = user_id
					defer session.Close()
					stmt, names := qb.Select("orders").Where(qb.Eq("user_id")).ToCql()
					q := gocqlx.Query(session.Query(stmt), names).BindMap(queryMap)
					if err := gocqlx.Select(&order, q.Query); err != nil {
						log.Println("ERROR in select query execution", err)
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
			formatter.JSON(w, http.StatusUnauthorized, "Bad Authentication")
		}
	}	
}

// API Delete an Order
func starbcuksDeleteOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("DELETE ORDERS BEGIN")
		token := req.Header.Get("Authorization")
		//token := "jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMGRiODFjYzAtY2FiMy00YzdlLWJjMjctODJjMDA1ZmE0NWMzIn0.sXx20MzYPM01lI7FYDSVTcgeJaeco36jWBIW5lVYuvI"
	
		userIdS, found := getUser(token)
		if found{			
			log.Println("Userid", userIdS)
			var pay Deletepay
			decoder := json.NewDecoder(req.Body)
			if err := decoder.Decode(&pay); err != nil {
				log.Println(err)
				formatter.JSON(w, http.StatusInternalServerError, "Something went bad with payload")
				return
			}
			log.Println("PID FOUND:",pay.Pid)
			pay_idS := pay.Pid 
			session, errs := cluster.CreateSession()
			if errs != nil{
				msg := "Error creating session"
				log.Println("ERROR session creation", errs)
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
							log.Println("ERROR in select count(*) query execution", err)
							formatter.JSON(w, http.StatusInternalServerError, err)
							return
						}
						if count == 0{
							log.Println("QUery returned 0 count")
							formatter.JSON(w, http.StatusInternalServerError, "Bad Payload")
							return
						}
						if err := session.Query(`delete from orders where user_id = ? and pay_id = ?`,user_id,pay_id ).Exec(); err != nil {
							log.Println("ERROR in delete query execution", err)
							formatter.JSON(w, http.StatusInternalServerError, err)
							return
						}
						log.Println("Sucessfully deleted")
						formatter.JSON(w, http.StatusAccepted, "Deleted")
						return
					}else{
						log.Println("ERROR in parsing Paydetail", err)
						formatter.JSON(w, http.StatusInternalServerError, err)
						return
					}
				}else{
					log.Println("ERROR in parsing User detail", err)
					formatter.JSON(w, http.StatusUnauthorized, "Unable to Authenticate")
					return
				}
			}
		}else{
			formatter.JSON(w, http.StatusUnauthorized, "Bad Authentication")
		}
	}
}	