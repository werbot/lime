package models

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// SMTP is ...
type SMTP struct {
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Encryption string `json:"encryption,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
}

// Validate is ...
func (v SMTP) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Host, is.Host),
		validation.Field(&v.Port, is.Port),
		validation.Field(&v.Username, validation.Length(3, 20)),
		validation.Field(&v.Password, validation.Length(3, 20)),
	)
}

// Mail is ...
type Mail struct {
	SenderName  string `json:"sender_name"`
	SenderEmail string `json:"sender_email"`
	SMTP        SMTP   `json:"smtp"`
}

// Validate is ...
func (v Mail) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.SenderName, validation.Length(2, 30)),
		validation.Field(&v.SenderEmail, is.Email),
		validation.Field(&v.SMTP),
	)
}

// Letter ...
type Letter struct {
	Subject string `json:"subject"`
	Text    string `json:"text"`
	Html    string `json:"html"`
}

// Validate is ...
func (v Letter) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Subject, validation.Length(5, 255)),
	)
}

// MessageMail ...
type MessageMail struct {
	To     string            `json:"to"`
	Letter Letter            `json:"letter"`
	Data   map[string]string `json:"data"`
}

// Validate is ...
func (v MessageMail) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.To, is.Email),
	)
}
