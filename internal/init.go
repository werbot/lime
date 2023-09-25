package app

import (
	"os"
	"path/filepath"

	"github.com/werbot/lime/pkg/crypto"
	"github.com/werbot/lime/pkg/fsutil"
)

const (
	configFile         = "./lime.toml"
	keyDir             = "./lime_keys"
	jwtPubKeyFile      = "jwt_public.key"
	jwtPrivKeyFile     = "jwt_private.key"
	licensePubKeyFile  = "license_public.key"
	licensePrivKeyFile = "license_private.key"
)

// Init is ...
func Init() error {
	if err := fsutil.MkDirs(0775, keyDir); err != nil {
		return err
	}
	GenJWTKeys()
	GenLicenseKeys()

	return nil
}

// GenConfigFile is ...
func GenConfigFile() error {
	config := DefaultConfig()
	if err := SaveConfig(config); err != nil {
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

	pubKeyPath := filepath.Join(keyDir, jwtPubKeyFile)
	privKeyPath := filepath.Join(keyDir, jwtPrivKeyFile)

	return generateAndSaveKeys(pubKeyPath, privKeyPath, jwtKey)
}

// GenLicenseKeys is ...
func GenLicenseKeys() error {
	licKey, err := crypto.GenerateEd25519()
	if err != nil {
		return err
	}

	pubKeyPath := filepath.Join(keyDir, licensePubKeyFile)
	privKeyPath := filepath.Join(keyDir, licensePrivKeyFile)

	return generateAndSaveKeys(pubKeyPath, privKeyPath, licKey)
}

func generateAndSaveKeys(pubKeyPath, privKeyPath string, key *crypto.Key) error {
	if err := os.Remove(pubKeyPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.WriteFile(pubKeyPath, key.PublicKeyToPEM(), 0666); err != nil {
		return err
	}

	if err := os.Remove(privKeyPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.WriteFile(privKeyPath, key.PrivateKeyToPEM(), 0666); err != nil {
		return err
	}

	return nil
}
