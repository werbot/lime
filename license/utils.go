package license

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

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

// KeyPair is a ...
type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

// KeyPairGenerate is a ...
func KeyPairGenerate() *KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &KeyPair{
		PublicKey:  base64.StdEncoding.EncodeToString(publicKey),
		PrivateKey: base64.StdEncoding.EncodeToString(privateKey),
	}
}
