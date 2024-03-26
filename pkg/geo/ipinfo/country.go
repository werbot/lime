package ipinfo

import (
	"net"

	"github.com/oschwald/maxminddb-golang"
)

// CountryRecord is ...
type CountryRecord struct {
	Country string `maxminddb:"country"`
}

// GetCountryCode is ...
func (cfg *Config) GetCountryCode(ip string) (string, error) {
	mmdb, err := maxminddb.Open(cfg.mmdbPath)
	if err != nil {
		return "", err
	}
	defer mmdb.Close()

	record := CountryRecord{}
	if err = mmdb.Lookup(net.ParseIP(ip), &record); err != nil {
		return "", err
	}
	return record.Country, nil
}
