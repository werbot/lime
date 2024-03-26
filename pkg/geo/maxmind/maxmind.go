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
	DBName     string `toml:"db-name,commented"`
	AccountID  int    `toml:"account-id,commented"`
	LicenseKey string `toml:"license-key,commented"`

	dbPath   string
	mmdbPath string
}

// DB is ...
func DB(dbPath string, cfg Config) *Config {
	cfg.dbPath = dbPath
	cfg.mmdbPath = filepath.Join(dbPath, cfg.DBName)
	return &cfg
}

// Check is ...
func (cfg *Config) Check() bool {
	return !fsutil.IsFile(cfg.mmdbPath) && cfg.AccountID > 0 && len(cfg.LicenseKey) > 30
}

// Download is ...
func (cfg *Config) Download() error {
	maxmindURL := fmt.Sprintf("https://%v:%s@download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz", cfg.AccountID, cfg.LicenseKey)
	pathArch := filepath.Join(cfg.dbPath, "GeoLite2-Country.tar.gz")

	// Download the file.
	if err := fsutil.Download(pathArch, maxmindURL); err != nil {
		return err
	}
	defer os.Remove(pathArch)

	// Extract the archive.
	if err := archive.ExtractTar(pathArch, cfg.dbPath); err != nil {
		return err
	}

	// Find the extracted database file.
	pathDB, err := filepath.Glob(filepath.Join(cfg.dbPath, "*", cfg.DBName))
	if err != nil {
		return err
	}
	if len(pathDB) == 0 {
		return fmt.Errorf("no DB file found in extracted path")
	}

	// Copy the database file to the desired location.
	if err := fsutil.CopyFile(pathDB[0], cfg.mmdbPath); err != nil {
		return err
	}

	defer fsutil.RemoveDir(filepath.Dir(pathDB[0]))

	return nil
}
