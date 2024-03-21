package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Term int

const (
	_ Term = iota
	Hour
	Day
	Week
	Month
	Year
)

type Currency int

const (
	_ Currency = iota
	CurrencyEUR
	CurrencyUSD
	CurrencyJPY
	CurrencyGBP
	CurrencyAUD
	CurrencyCAD
	CurrencyCHF
	CurrencyCNY
	CurrencySEK
)

// Patterns is ...
type Patterns struct {
	Total    int        `json:"total"`
	Patterns []*Pattern `json:"patterns,omitempty"`
}

// Pattern is a ...
type Pattern struct {
	Core
	Name     string    `json:"name"`
	Limit    Metadata  `json:"limit,omitempty"`    // license restrictions
	Term     Term      `json:"term"`               // license validity period
	Price    int       `json:"price"`              // price in Stripe format
	Currency Currency  `json:"currency"`           // currency
	Check    Metadata  `json:"check,omitempty"`    // what will be checked for license verification?
	Private  bool      `json:"private,omitempty"`  // the pattern is available to the administrator only.
	Status   bool      `json:"status,omitempty"`   // pattern activity
	Licenses *Licenses `json:"licenses,omitempty"` // licenses
}

// Validate is ...
func (v Pattern) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Length(5, 128)),
		validation.Field(&v.Term, validation.In("d", "w", "m", "y")), // d-day, w-week, m-month, y-year
		validation.Field(&v.Price, validation.Required, validation.Min(0)),
	)
}
