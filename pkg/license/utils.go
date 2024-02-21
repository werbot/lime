package license

import (
	"encoding/pem"

	"golang.org/x/crypto/ed25519"
)

// DecodePublicKey is a ...
func DecodePublicKey(data []byte) ed25519.PublicKey {
	block, _ := pem.Decode(data)
	return ed25519.PublicKey(block.Bytes)
}

// DecodePrivateKey is a ...
func DecodePrivateKey(data []byte) ed25519.PrivateKey {
	block, _ := pem.Decode(data)
	return ed25519.PrivateKey(block.Bytes)
}
