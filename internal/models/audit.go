package models

import (
	"time"

	"github.com/werbot/lime/pkg/webutil"
)

type Section int

const (
	_ Section = iota
	SectionSystem
	SectionSetting
	SectionCustomer
	SectionPattern
	SectionLicense
	SectionPayment
)

type AuditAction int

const (
	_ AuditAction = iota
	OnSendMail
	OnSignIn
	OnSignOut
	OnAdd
	OnUpdate
	OnDelete
	OnClone
)

// Audits is ...
type Audits struct {
	Total  int      `json:"total"`
	Audits []*Audit `json:"audits"`
}

// Audit is ...
type Audit struct {
	ID       string           `json:"id"`
	Section  Section          `json:"section"`
	Customer Customer         `json:"customer"`
	Action   AuditAction      `json:"action"`
	Metadata webutil.MetaInfo `json:"metadata"`
	Created  time.Time        `json:"created"`
}
