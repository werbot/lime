package ipinfo

import (
	"fmt"
	"path/filepath"

	"github.com/werbot/lime/pkg/fsutil"
)

// Config is ...
type Config struct {
	DBName string `toml:"db-name"`
	Token  string `toml:"token"`
}

// Download is ...
func (cfg Config) Download(dbPath string) error {
	maxmindURL := fmt.Sprintf("https://ipinfo.io/data/free/country.mmdb?token=%s", cfg.Token)
	mmdbPath := filepath.Join(dbPath, cfg.DBName)

	// Download the file.
	if err := fsutil.Download(mmdbPath, maxmindURL); err != nil {
		return err
	}

	return nil
}
