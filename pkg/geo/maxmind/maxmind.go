package maxmind

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/werbot/lime/pkg/archive"
	"github.com/werbot/lime/pkg/fsutil"
)

// Config is ...
type Config struct {
	DBName     string `toml:"db-name"`
	AccountID  int    `toml:"account-id"`
	LicenseKey string `toml:"license-key"`
}

// Download is ...
func (cfg Config) Download(dbPath string) error {
	maxmindURL := fmt.Sprintf("https://%v:%s@download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz", cfg.AccountID, cfg.LicenseKey)
	pathArch := filepath.Join(dbPath, "GeoLite2-Country.tar.gz")

	// Download the file.
	if err := fsutil.Download(pathArch, maxmindURL); err != nil {
		return err
	}
	defer os.Remove(pathArch)

	// Extract the archive.
	if err := archive.ExtractTar(pathArch, dbPath); err != nil {
		return err
	}

	// Find the extracted database file.
	pathDB, err := filepath.Glob(filepath.Join(dbPath, "*", cfg.DBName))
	if err != nil {
		return err
	}
	if len(pathDB) == 0 {
		return fmt.Errorf("no DB file found in extracted path")
	}

	// Copy the database file to the desired location.
	mmdbPath := filepath.Join(dbPath, cfg.DBName)
	if err := fsutil.CopyFile(pathDB[0], mmdbPath); err != nil {
		return err
	}

	defer fsutil.RemoveDir(filepath.Dir(pathDB[0]))

	return nil
}
