package geo

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/oschwald/maxminddb-golang"

	"github.com/werbot/lime/pkg/archive"
	"github.com/werbot/lime/pkg/fsutil"
)

// Maxmind is ...
type Maxmind struct {
	DBPath     string `toml:"db-path"`
	AccountID  int    `toml:"account-id"`
	LicenseKey string `toml:"license-key"`
}

// Download is ...
func Download(cfg Maxmind) error {
	maxmindURL := fmt.Sprintf("https://%v:%s@download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz", cfg.AccountID, cfg.LicenseKey)

	// Construct the archive path using filepath.Join for cross-platform compatibility.
	pathArch := filepath.Join(cfg.DBPath, "GeoLite2-Country.tar.gz")

	// Download the file.
	if err := fsutil.Download(pathArch, maxmindURL); err != nil {
		return err
	}
	defer os.Remove(pathArch)

	// Extract the archive.
	if err := archive.ExtractTar(pathArch, cfg.DBPath); err != nil {
		return err
	}

	// Find the extracted database file.
	pathDB, err := filepath.Glob(filepath.Join(cfg.DBPath, "*", "GeoLite2-Country.mmdb"))
	if err != nil {
		return err
	}
	if len(pathDB) == 0 {
		return fmt.Errorf("no DB file found in extracted path")
	}

	// Copy the database file to the desired location.
	mmdbPath := filepath.Join(cfg.DBPath, "GeoLite2-Country.mmdb")
	if err := fsutil.CopyFile(pathDB[0], mmdbPath); err != nil {
		return err
	}

	defer fsutil.RemoveDir(filepath.Dir(pathDB[0]))

	return nil
}

// Country is ...
func Country(dbPath, ip string) (string, error) {
	db, err := maxminddb.Open(dbPath)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	err = db.Lookup(net.ParseIP(ip), &record)
	if err != nil {
		log.Panic(err)
	}

	return record.Country.ISOCode, nil
}
