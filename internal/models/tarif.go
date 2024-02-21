package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Tariff is a ...
type Tariff struct {
	Core
	Name     string     `json:"name"`
	Amount   int        `json:"amount"`
	Metadata []Metadata `json:"metadata,omitempty"`
}

// Validate is ...
func (v Tariff) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Length(3, 50)),
		validation.Field(&v.Amount, validation.Required, validation.Min(0)),
		validation.Field(&v.Metadata),
	)
}

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
