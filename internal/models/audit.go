package models

import "time"

type Section int

const (
	_ Section = iota
	SectionSetting
	SectionCustomer
	SectionPattern
	SectionLicense
)

type Action int

const (
	_ Action = iota
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
	ID        string    `json:"id"`
	Section   Section   `json:"section"`
	SectionID string    `json:"section_id"`
	Action    Action    `json:"action"`
	Metadata  Metadata  `json:"metadata"`
	Created   time.Time `json:"created"`
}
