package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Customers is ...
type Customers struct {
	Total     int        `json:"total"`
	Customers []Customer `json:"customers"`
}

// Customer is a ...
type Customer struct {
	Core
	Email    string     `json:"email"`
	Status   bool       `json:"status"`
	Payments *[]Payment `json:"payments,omitempty"`
}

// Validate is ...
func (v Customer) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required, is.Email),
		validation.Field(&v.Payments),
	)
}
