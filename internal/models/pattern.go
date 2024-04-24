package models

import (
	"time"

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

// ToDuration converts a Term to its equivalent time.Duration.
func (t Term) ToDuration() time.Duration {
	switch t {
	case Hour:
		return time.Hour
	case Day:
		return 24 * time.Hour
	case Week:
		return 7 * 24 * time.Hour
	case Month:
		return 30 * 24 * time.Hour // Approximation, as months vary in length.
	case Year:
		return 365 * 24 * time.Hour // Approximation, ignoring leap years.
	default:
		return 0 // Undefined term.
	}
}

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
	Limit    *Metadata `json:"limit,omitempty"`    // license restrictions
	Term     *Term     `json:"term,omitempty"`     // license validity period
	Price    *int      `json:"price,omitempty"`    // price in Stripe format
	Currency *Currency `json:"currency,omitempty"` // currency
	Check    *Metadata `json:"check,omitempty"`    // what will be checked for license verification?
	Private  bool      `json:"private,omitempty"`  // the pattern is available to the administrator only.
	Status   bool      `json:"status,omitempty"`   // pattern activity
	Licenses *Licenses `json:"licenses,omitempty"` // licenses
}

// Validate is ...
func (v Pattern) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Length(5, 128)),
		validation.Field(&v.Price, validation.Required, validation.Min(0)),
	)
}
