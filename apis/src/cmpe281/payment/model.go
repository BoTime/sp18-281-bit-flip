package main

import (
	"github.com/gocql/gocql"
)

// PaymentDetails model
type PaymentDetails struct {
	UserId         gocql.UUID     `json:"-" cql:"user_id"`
	PaymentId      gocql.UUID     `json:"payment_id" cql:"payment_id"`
	Amount         string         `json:"amount" cql:"amount"`
	Status         string         `json:"status" cql:"status"`
	CardDetails    CardDetails    `json:"card_details" cql:"card_details"`
	BillingDetails BillingDetails `json:"billing_details" cql:"billing_details"`
}

// CardDetails model
type CardDetails struct {
	Number   string `json:"number" cql:"number"`
	Cvv      string `json:"cvv" cql:"cvv"`
	ExpMonth string `json:"exp_month" cql:"exp_month"`
	ExpYear  string `json:"exp_year" cql:"exp_year"`
}

// BillingDetails model
type BillingDetails struct {
	FirstName string `json:"first_name" cql:"first_name"`
	LastName  string `json:"last_name" cql:"last_name"`
	Line1     string `json:"line1" cql:"line1"`
	Line2     string `json:"line2" cql:"line2"`
	City      string `json:"city" cql:"city"`
	State     string `json:"state" cql:"state"`
	ZipCode   string `json:"zip_code" cql:"zip_code"`
}

// ListPaymentsResult model
type ListPaymentsResult struct {
	Payments      []PaymentDetails `json:"payments"`
	NextPageToken *gocql.UUID      `json:"next_page_token"`
}
