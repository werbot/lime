package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PaymentProvider int

const (
	_ PaymentProvider = iota
	NONE
	STRIPE
)

type PaymentStatus int

const (
	_ PaymentStatus = iota
	PAID
	UNPAID
	PROCESSED
	CANCELED
	FAILED
)

// Payments is ...
type Payments struct {
	Total    int        `json:"total"`
	Payments []*Payment `json:"payments,omitempty"`
}

// Payment is a ...
type Payment struct {
	Core
	Customer    *Customer    `json:"customer"`
	Pattern     *Pattern     `json:"pattern"`
	Transaction *Transaction `json:"transaction,omitempty"`
}

// Validate is ...
func (v Payment) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Customer),
		validation.Field(&v.Pattern),
		validation.Field(&v.Transaction),
	)
}

// TransactionTransactionTransaction is ...
type Transaction struct {
	Provider PaymentProvider `json:"provider"`
	Status   PaymentStatus   `json:"status"`
	Meta     Metadata        `json:"meta,omitempty"`
}
