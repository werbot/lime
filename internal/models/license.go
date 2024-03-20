package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Licenses is ...
type Licenses struct {
	Total    int       `json:"total"`
	Licenses []License `json:"licenses"`
}

// License is a ...
type License struct {
	Core
	Customer *Customer `json:"customer,omitempty"`
	Pattern  *Pattern  `json:"pattern,omitempty"`
	License  []byte    `json:"license"`
	Hash     string    `json:"hash,omitempty"`
	Status   bool      `json:"status"`
}

// Validate is ...
func (v License) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Pattern),
	)
}
