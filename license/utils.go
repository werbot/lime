package license

import (
	"encoding/base64"

	"golang.org/x/crypto/ed25519"
)

// DecodePublicKey is a ...
func DecodePublicKey(data []byte) (ed25519.PublicKey, error) {
	decoded, err := decode(data)
	if err != nil {
		return nil, err
	}
	return ed25519.PublicKey(decoded), nil
}

// DecodePrivateKey is a ...
func DecodePrivateKey(data []byte) (ed25519.PrivateKey, error) {
	decoded, err := decode(data)
	if err != nil {
		return nil, err
	}
	return ed25519.PrivateKey(decoded), nil
}

func decode(b []byte) ([]byte, error) {
	enc := base64.StdEncoding
	buf := make([]byte, enc.DecodedLen(len(b)))
	n, err := enc.Decode(buf, b)
	return buf[:n], err
}
