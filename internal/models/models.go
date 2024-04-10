package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Core is ...
type Core struct {
	ID      string     `json:"id"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

// Validate is ...
func (v Core) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(21, 21)),
	)
}

// Metadata is ...
type Metadata map[string]any
