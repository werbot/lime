package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/geo"
	"github.com/werbot/lime/pkg/geo/geoopen"
	"github.com/werbot/lime/pkg/geo/ipinfo"
	"github.com/werbot/lime/pkg/geo/maxmind"
	"github.com/werbot/lime/pkg/storage"
	"github.com/werbot/lime/pkg/storage/postgres"
	"github.com/werbot/lime/pkg/storage/sqlite"
)

const (
	ConfigFile = "./lime.toml"
)

var cfg *Config

// Config is ...
type Config struct {
	HTTPAddr    string           `toml:"http-addr" comment:"Ports <= 1024 are privileged ports. You can't use them unless you're root or have the explicit\npermission to use them. See this answer for an explanation or wikipedia or something you trust more.\nsudo setcap 'cap_net_bind_service=+ep' /opt/yourGoBinary"`
	DevMode     bool             `toml:"dev-mode" comment:"Active develop mode"`
	Admin       Admin            `toml:"admin" comment:"Admin section"`
	Keys        Keys             `toml:"keys" comment:"Keys section"`
	GeoDatabase geo.Database     `toml:"geo-database" comment:"Geo database section"`
	Database    storage.Database `toml:"database" comment:"Database section"`
}

// Admin is ...
type Admin struct {
	Email    string `toml:"email"`
	Password string `toml:"password"`
}

// Keys is ...
type Keys struct {
	KeyDir  string  `toml:"key-dir"`
	JWT     JWT     `toml:"jwt"`
	License License `toml:"license"`
}

// JWT is ...
type JWT struct {
	PublicKey  string `toml:"public-key"`
	PrivateKey string `toml:"private-key"`
	Expire     string `toml:"expire"`
}

// License is ...
type License struct {
	PublicKey  string `toml:"public-key"`
	PrivateKey string `toml:"private-key"`
}

// DefaultConfig is ...
func DefaultConfig() *Config {
	return &Config{
		DevMode:  false,
		HTTPAddr: "0.0.0.0:8088",
		Admin: Admin{
			Email:    "admin@mail.com",
			Password: "Pass123",
		},
		Keys: Keys{
			KeyDir: "./lime_keys",
			JWT: JWT{
				PublicKey:  "jwt_public.key",
				PrivateKey: "jwt_private.key",
				Expire:     "10m",
			},
			License: License{
				PublicKey:  "license_public.key",
				PrivateKey: "license_private.key",
			},
		},
		GeoDatabase: geo.Database{
			DBPath:  "./lime_geo",
			Storage: geo.GeoOpen,
			GeoOpen: geoopen.Config{
				DBName: "GeoOpen-Country.mmdb",
			},
			Maxmind: maxmind.Config{
				DBName:     "GeoLite2-Country.mmdb",
				AccountID:  0,
				LicenseKey: "",
			},
			Ipinfo: ipinfo.Config{
				DBName: "IPinfo-Country.mmdb",
				Token:  "",
			},
		},
		Database: storage.Database{
			Storage: storage.Sqlite,
			Sqlite: sqlite.Config{
				DBPath: "./lime_base/data.db",
			},
			Postgres: postgres.Config{
				User:            "user",
				Password:        "password",
				Host:            "localhost:5432",
				Database:        "lime",
				MaxConn:         50,
				MaxIdleConn:     10,
				MaxLifetimeConn: 300,
			},
		},
	}
}

// LoadConfig is ...
func LoadConfig() error {
	config := DefaultConfig()

	if fsutil.IsFile(ConfigFile) {
		file, err := os.ReadFile(ConfigFile)
		if err != nil {
			return err
		}

		if err := toml.Unmarshal(file, &config); err != nil {
			return err
		}
	}

	cfg = config
	return nil
}

// SaveConfig is ...
func SaveConfig(config *Config) error {
	byteConfig, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	if err := os.WriteFile(ConfigFile, byteConfig, 0o666); err != nil {
		return err
	}

	return nil
}

// Data is ...
func Data() *Config {
	if cfg == nil {
		cfg = &Config{}
	}
	return cfg
}
