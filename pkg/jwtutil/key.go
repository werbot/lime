package jwtutil

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/werbot/lime/pkg/fsutil"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

// LoadKeys is ...
func LoadKeys(pubKeyPath, privKeyPath string) error {
	// Parse RSA Public Key from PEM encoded bytes
	pubBytes := fsutil.MustReadFile(pubKeyPath)
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %w", err)
	}

	// Parse RSA Private Key from PEM encoded bytes
	privBytes := fsutil.MustReadFile(privKeyPath)
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey = pubKey
	privateKey = privKey

	return nil
}
