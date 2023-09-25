package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type Key struct {
	publicKey  []byte
	privateKey []byte
}

func GenerateRSA() (*Key, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &Key{
		publicKey:  x509.MarshalPKCS1PublicKey(&privKey.PublicKey),
		privateKey: x509.MarshalPKCS1PrivateKey(privKey),
	}, nil
}

func GenerateEd25519() (*Key, error) {
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}

	return &Key{
		publicKey:  pubKey,
		privateKey: privKey,
	}, nil
}

func (k *Key) PublicKeyToPEM() []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: k.publicKey})
}

func (k *Key) PrivateKeyToPEM() []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: k.privateKey})
}
