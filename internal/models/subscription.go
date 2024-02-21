package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Subscription is a ...
type Subscription struct {
	Core
	StripeID   string    `json:"stripe_id"`
	CustomerID string    `json:"customer_id"`
	TariffID   string    `json:"tariff_id"`
	Status     bool      `json:"status"`
	Customer   Customer  `json:"customer"`
	Tariff     Tariff    `json:"tariff"`
	IssuedAt   time.Time `json:"issued_at"`  // Issued At
	ExpiresAt  time.Time `json:"expires_at"` // Expires At
	Licenses   []License `json:"licenses"`
}

// Validate is ...
func (v Subscription) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.CustomerID, validation.Length(21, 21)),
		validation.Field(&v.TariffID, validation.Length(21, 21)),
		validation.Field(&v.Customer),
		validation.Field(&v.Tariff),
		validation.Field(&v.Licenses),
	)
}
