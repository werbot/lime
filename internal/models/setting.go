package models

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// SettingName is ...
type SettingName struct {
	ID    string `json:"id,omitempty"`
	Key   string `json:"key"`
	Value any    `json:"value,omitempty"`
}

// Validate is ...
func (v SettingName) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Key, validation.Required),
	)
}

// Site is ...
type Site struct {
	Domain       string `json:"domain"`
	Name         string `json:"name"`
	Signature    string `json:"signature"`
	EmailSupport string `json:"email_support"`
}

// Validate is ...
func (v Site) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Domain, is.Host),
		validation.Field(&v.Name, validation.Length(5, 128)),
		validation.Field(&v.Signature, validation.Length(5, 128)),
		validation.Field(&v.EmailSupport, is.Email),
	)
}

// SMTP is ...
type SMTP struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Encryption int    `json:"encryption"`
	Username   string `json:"username"`
	Password   string `json:"password"`
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

// MailMessage ...
type MailMessage struct {
	To     string            `json:"to"`
	Letter Letter            `json:"letter"`
	Data   map[string]string `json:"data"`
}

// Validate is ...
func (v MailMessage) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.To, is.Email),
	)
}
