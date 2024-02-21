package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// User is a ...
type User struct {
	Core
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
	Role     string `json:"role"`
}

// Validate is ...
func (v User) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Login, validation.Length(3, 50)),
		validation.Field(&v.Email, validation.Required, is.Email),
	)
}
