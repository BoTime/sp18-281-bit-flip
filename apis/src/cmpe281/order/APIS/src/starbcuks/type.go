package main

import 	"github.com/gocql/gocql"

// Ordered Product detail
type OrderDetails struct {
	Item    string    `json:"item" cql:"item"`
	Quantity	string    `json:"qty" cql:"qty"`
	Size string `json:"size" cql:"size"`
}

// Delete Order 
type Deletepay struct {
	Pid    string    `json:"pid"`
}

// Store All Payment related data
type Payments struct {
	Amount   string    `json:"amount" cql:"amount"`
	Bill	Billing    `json:"billing_details" cql:"billing_details"`
	CardDetails Card `json:"card_details" cql:"card_details"`
}

type Billing struct {
	First string `json:"first_name" cql:"first_name"`
    Second string `json:"last_name" cql:"last_name"`
	Add1 string `json:"line1" cql:"line1"`
    Add2 string `json:"line2" cql:"line2"`
	City string `json:"city" cql:"city"`
    State string `json:"state" cql:"state"`
    Pin string `json:"zip_code" cql:"zip_code"`
}
type Card struct {
	Number string `json:"number" cql:"number"`
    Exp_month string `json:"exp_month" cql:"exp_month"`
    Exp_yr string `json:"exp_year" cql:"exp_year"`
	Cvv string `json:"cvv" cql:"cvv"`
}

// Store complete order detail for Gets
type Order struct { 
  UserId  gocql.UUID  `json:"-" cql:"user_id"`
  PayId  gocql.UUID  `json:"-" cql:"pay_id"`
  Status string `json:"status" cql:"status"`
  Store  string    `json:"store" cql:"store"`
  Product []OrderDetails `json:"product" cql:"product"`
  Payment Payments `json:"payment" cql:"payment"`
  }

// Map database for Responses for Gets
type GetOrder struct { 
  UserId  gocql.UUID  `json:"-" cql:"user_id"`
  PayId  gocql.UUID  `json:"pay_id" cql:"pay_id"`
  Status string `json:"status" cql:"status"`
  Store  string    `json:"store" cql:"store"`
  Product []OrderDetails `json:"product" cql:"product"`
  Payment Payments `json:"-" cql:"payment"`
}

type GetPayments struct {
	Amount   string    `json:"amount"`
	Bill	Billing    `json:"billing_details"`
	CardDetails Card `json:"card_details"`
	PayID gocql.UUID `json:"payment_id"`
	Status string `json:"status"`
}
type GetInventory struct { 
  InvId  gocql.UUID  `json:"id"`
  Status string `json:"status"`
  Expires  int    `json:"expires"`
  Product []OrderDetails `json:"products"`
}