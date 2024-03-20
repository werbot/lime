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

/*
// Metadata is ...
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Validate is ...
func (v Metadata) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Key, validation.Required, validation.Length(1, 20)),
		validation.Field(&v.Value, validation.Required, validation.Min(0)),
	)
}
*/

type Metadata map[string]any
