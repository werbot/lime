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

	// generate new ed25519 key and replaces !!!
	privateKey = []byte("5GvXN6OIrsgF3/ehJ17HvRPrrbNTLw/gtAmy4X5bKlH9rmXwQgFSVLt//nMsl0qFG28pjc1IN7PhgH01Z+QCTQ==")
	publicKey  = []byte("/a5l8EIBUlS7f/5zLJdKhRtvKY3NSDez4YB9NWfkAk0=")
)

// License is a ...
type License struct {
	Iss string          `json:"iss,omitempty"` // Issued By
	Cus string          `json:"cus,omitempty"` // Customer ID
	Sub uint32          `json:"sub,omitempty"` // Subscriber ID
	Typ string          `json:"typ,omitempty"` // License Type
	Lim Limits          `json:"lim,omitempty"` // License Limit (e.g. Site)
	Iat time.Time       `json:"iat,omitempty"` // Issued At
	Exp time.Time       `json:"exp,omitempty"` // Expires At
	Dat json.RawMessage `json:"dat,omitempty"` // Metadata
}

// Limits is a ...
type Limits struct {
	Servers   int `json:"servers"`
	Companies int `json:"companies"`
	Users     int `json:"users"`
}

// Expired is a ...
func (l *License) Expired() bool {
	return l.Exp.IsZero() == false && time.Now().After(l.Exp)
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

// GetPrivateKey is a ...
func GetPrivateKey() ed25519.PrivateKey {
	key, _ := DecodePrivateKey(privateKey)
	return key
}

// GetPublicKey is a ...
func GetPublicKey() ed25519.PublicKey {
	key, _ := DecodePublicKey(publicKey)
	return key
}
