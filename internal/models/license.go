package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// License is a ...
type License struct {
	Core
	SubscriptionID string       `json:"subscription_id"`
	Mac            string       `json:"mac"`
	License        []byte       `json:"license"`
	Hash           string       `json:"hash"`
	Status         bool         `json:"status"`
	Subscription   Subscription `json:"subscription"`
}

// Validate is ...
func (v License) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.SubscriptionID, validation.Length(21, 21)),
		validation.Field(&v.Mac, is.MAC),
		validation.Field(&v.Subscription),
	)
}
