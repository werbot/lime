package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Licenses is ...
type Licenses struct {
	Total    int        `json:"total"`
	Licenses []*License `json:"licenses,omitempty"`
}

// License is a ...
type License struct {
	Core
	Payment *Payment `json:"payment,omitempty"`
	License *[]byte  `json:"license,omitempty"`
	Hash    string   `json:"hash,omitempty"`
	Status  bool     `json:"status"`
}

// Validate is ...
func (v License) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Payment),
	)
}
