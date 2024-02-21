package app

import (
	"os"
	"path/filepath"

	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/pkg/crypto"
)

// GenConfigFile is ...
func GenConfigFile() error {
	cfg := config.DefaultConfig()
	if err := config.SaveConfig(cfg); err != nil {
		return err
	}
	return nil
}

// GenJWTKeys is ...
func GenJWTKeys() error {
	jwtKey, err := crypto.GenerateRSA()
	if err != nil {
		return err
	}

	pubKeyPath := filepath.Join(config.KeyDir, config.JWTPubKeyFile)
	privKeyPath := filepath.Join(config.KeyDir, config.JWTPrivKeyFile)
	return generateAndSaveKeys(pubKeyPath, privKeyPath, jwtKey)
}

// GenLicenseKeys is ...
func GenLicenseKeys() error {
	licKey, err := crypto.GenerateEd25519()
	if err != nil {
		return err
	}

	pubKeyPath := filepath.Join(config.KeyDir, config.LicensePubKeyFile)
	privKeyPath := filepath.Join(config.KeyDir, config.LicensePrivKeyFile)
	return generateAndSaveKeys(pubKeyPath, privKeyPath, licKey)
}

func generateAndSaveKeys(pubKeyPath, privKeyPath string, key *crypto.Key) error {
	if err := os.Remove(pubKeyPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.WriteFile(pubKeyPath, key.PublicKeyToPEM(), 0o644); err != nil {
		return err
	}

	if err := os.Remove(privKeyPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.WriteFile(privKeyPath, key.PrivateKeyToPEM(), 0o600); err != nil {
		return err
	}

	return nil
}
