package models

import "time"

type Section int

const (
	SectionSystem Section = iota
	SectionSetting
	SectionCustomer
	SectionPattern
	SectionLicense
)

type AuditAction int

const (
	OnSendMail AuditAction = iota
	OnSignIn
	OnSignOut
	OnAdd
	OnUpdate
	OnDelete
)

// Audits is ...
type Audits struct {
	Total  int     `json:"total"`
	Audits []Audit `json:"audits"`
}

// Audit is ...
type Audit struct {
	ID       string      `json:"id"`
	Section  Section     `json:"section"`
	Customer Customer    `json:"customer"`
	Action   AuditAction `json:"action"`
	Metadata Metadata    `json:"metadata"`
	Created  time.Time   `json:"created"`
}
