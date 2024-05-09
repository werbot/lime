package license

import (
	"encoding/json"
	"time"

	"golang.org/x/crypto/ed25519"
)

// Private is ...
type Private struct {
	key     ed25519.PrivateKey
	License License
}

// Public is ...
type Public struct {
	key     ed25519.PublicKey
	License License
}

// License is a ...
type License struct {
	IssuedBy     string          `json:"iss,omitempty"`
	CustomerID   string          `json:"cus,omitempty"`
	SubscriberID string          `json:"sub,omitempty"`
	Type         string          `json:"typ,omitempty"`
	Limit        Limits          `json:"lim,omitempty"`
	IssuedAt     time.Time       `json:"iat,omitempty"`
	ExpiresAt    time.Time       `json:"exp,omitempty"`
	Metadata     json.RawMessage `json:"dat,omitempty"`
}

// Limits is ...
type Limits struct {
	Limits []Limit `json:"limits"`
}

// Limit is ...
type Limit struct {
	Key   string   `json:"key"`
	Value int      `json:"value,omitempty"`
	List  []string `json:"list,omitempty"`
}
