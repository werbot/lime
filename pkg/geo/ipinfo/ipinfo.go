package ipinfo

import (
	"fmt"
	"path/filepath"

	"github.com/werbot/lime/pkg/fsutil"
)

// Config is ...
type Config struct {
	DBName string `toml:"db-name,commented"`
	Token  string `toml:"token,commented"`

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
	return !fsutil.IsFile(cfg.mmdbPath) && len(cfg.Token) > 10
}

// Download is ...
func (cfg *Config) Download() error {
	maxmindURL := fmt.Sprintf("https://ipinfo.io/data/free/country.mmdb?token=%s", cfg.Token)

	// Download the file.
	if err := fsutil.Download(cfg.mmdbPath, maxmindURL); err != nil {
		return err
	}

	return nil
}
