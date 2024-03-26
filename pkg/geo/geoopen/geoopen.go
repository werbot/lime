package geoopen

import (
	"path/filepath"

	"github.com/werbot/lime/pkg/fsutil"
)

// Config is ...
type Config struct {
	DBName string `toml:"db-name"`

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
func (cfg Config) Check() bool {
	return !fsutil.IsFile(cfg.mmdbPath)
}

// Download is ...
func (cfg Config) Download() error {
	maxmindURL := "https://cra.circl.lu/opendata/geo-open/mmdb-country/latest.mmdb"

	// Download the file.
	if err := fsutil.Download(cfg.mmdbPath, maxmindURL); err != nil {
		return err
	}

	return nil
}
