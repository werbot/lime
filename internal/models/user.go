package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// User is a ...
type User struct {
	Core
	Email  string `json:"email"`
	Status bool   `json:"status"`
}

// Validate is ...
func (v User) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required, is.Email),
	)
}
