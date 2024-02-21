package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Core is ...
type Core struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

// Validate is ...
func (v Core) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(21, 21)),
	)
}
