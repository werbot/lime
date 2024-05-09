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
	// ErrInvalidLicense  is ...
	ErrInvalidLicense = errors.New("invalid license")

	// ErrMalformedLicense is a ...
	ErrMalformedLicense = errors.New("malformed license")

	// ErrPublicKeyNotSet is ...
	ErrPublicKeyNotSet = errors.New("public key is not set")
)

// DecodePrivateKey is decode private key from base64
func DecodePrivateKey(privateKey []byte) (*Private, error) {
	decodedPrivateKey, err := decodeKey(privateKey)
	if err != nil {
		return nil, err
	}

	return &Private{
		key: ed25519.PrivateKey(decodedPrivateKey),
	}, nil
}

// Encode is a generate new license
func (l *Private) Encode() ([]byte, error) {
	if l.key == nil {
		return nil, ErrPublicKeyNotSet
	}

	msg, err := json.Marshal(l.License)
	if err != nil {
		return nil, err
	}

	sig := ed25519.Sign(l.key, msg)
	buf := &bytes.Buffer{}
	buf.Write(sig)
	buf.Write(msg)

	block := &pem.Block{
		Type:  "LICENSE KEY",
		Bytes: buf.Bytes(),
	}
	return pem.EncodeToMemory(block), nil
}

// Expired is a ...
func (l *Private) Expired() bool {
	return !l.License.ExpiresAt.IsZero() && time.Now().After(l.License.ExpiresAt)
}

// DecodePublicKey and decode public key from base64
func DecodePublicKey(publicKey []byte) (*Public, error) {
	decodedPublicKey, err := decodeKey(publicKey)
	if err != nil {
		return nil, err
	}

	return &Public{
		key: ed25519.PublicKey(decodedPublicKey),
	}, nil
}

// Decode license file
func (l *Public) Decode(data []byte) (*Public, error) {
	if l.key == nil {
		return nil, ErrPublicKeyNotSet
	}

	block, _ := pem.Decode(data)
	if block == nil || len(block.Bytes) < ed25519.SignatureSize {
		return nil, ErrMalformedLicense
	}

	sig := block.Bytes[:ed25519.SignatureSize]
	msg := block.Bytes[ed25519.SignatureSize:]

	verified := ed25519.Verify(l.key, msg, sig)
	if !verified {
		return nil, ErrInvalidLicense
	}
	out := &License{}
	err := json.Unmarshal(msg, out)
	l.License = *out
	return l, err
}

// Expired is a ...
func (l *Public) Expired() bool {
	return !l.License.ExpiresAt.IsZero() && time.Now().After(l.License.ExpiresAt)
}

// Info is ...
func (l *Public) Info() License {
	return l.License
}
