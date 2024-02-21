package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Customer is a ...
type Customer struct {
	Core
	UserID        string         `json:"user_id"`
	Status        bool           `json:"status"`
	Subscriptions []Subscription `json:"subscriptions"`
}

// Validate is ...
func (v Customer) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.UserID, validation.Length(21, 21)),
		validation.Field(&v.Subscriptions),
	)
}
