package license

import (
	"bytes"
	"encoding/json"
	"encoding/pem"
	"errors"
	"time"

	"golang.org/x/crypto/ed25519"
)

var (
	// ErrInvalidSignature is a ...
	ErrInvalidSignature = errors.New("Invalid signature")

	// ErrMalformedLicense is a ...
	ErrMalformedLicense = errors.New("Malformed License")
)

// License is a ...
type License struct {
	IssuedBy     string          `json:"iss,omitempty"`
	CustomerID   string          `json:"cus,omitempty"`
	SubscriberID uint32          `json:"sub,omitempty"`
	Type         string          `json:"typ,omitempty"`
	Limit        []Limits        `json:"lim,omitempty"`
	IssuedAt     time.Time       `json:"iat,omitempty"`
	ExpiresAt    time.Time       `json:"exp,omitempty"`
	Metadata     json.RawMessage `json:"dat,omitempty"`
}

// Limits is ...
type Limits struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Expired is a ...
func (l *License) Expired() bool {
	return !l.ExpiresAt.IsZero() && time.Now().After(l.ExpiresAt)
}

// Encode is a ...
func (l *License) Encode(privateKey ed25519.PrivateKey) ([]byte, error) {
	msg, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	sig := ed25519.Sign(privateKey, msg)
	buf := new(bytes.Buffer)
	buf.Write(sig)
	buf.Write(msg)

	block := &pem.Block{
		Type:  "LICENSE KEY",
		Bytes: buf.Bytes(),
	}
	return pem.EncodeToMemory(block), nil
}

// Decode is a ...
func Decode(data []byte, publicKey ed25519.PublicKey) (*License, error) {
	block, _ := pem.Decode(data)
	if block == nil || len(block.Bytes) < ed25519.SignatureSize {
		return nil, ErrMalformedLicense
	}

	sig := block.Bytes[:ed25519.SignatureSize]
	msg := block.Bytes[ed25519.SignatureSize:]

	verified := ed25519.Verify(publicKey, msg, sig)
	if !verified {
		return nil, ErrInvalidSignature
	}
	out := new(License)
	err := json.Unmarshal(msg, out)
	return out, err
}
