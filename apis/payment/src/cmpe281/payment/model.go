package main

import (
	"github.com/gocql/gocql"
)

// PaymentDetails model
type PaymentDetails struct {
	UserId    gocql.UUID `db:"user_id" json:"-"`
	PaymentId gocql.UUID `json:"payment_id"`
	Amount    float64    `json:"amount"`
	Status    string     `json:"status"`
	CardDetails CardDetails `json:"card_details"`
	BillingDetails BillingDetails `json:"billing_details"`
}

// CardDetails model
type CardDetails struct {
	Number string `json:"number"`
	Expiration string `json:"expiration"`
}

// BillingDetails model
type BillingDetails struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zipcode"`
}

// ListPaymentsResult model
type ListPaymentsResult struct {
	Payments      []PaymentDetails `json:"payments"`
	NextPageToken *gocql.UUID       `json:"next_page_token"`
}