package app

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/storage"
	"github.com/werbot/lime/pkg/storage/postgres"
	"github.com/werbot/lime/pkg/storage/sqlite"
)

type Config struct {
	HTTPAddr string           `toml:"http-addr" comment:"Ports <= 1024 are privileged ports. You can't use them unless you're root or have the explicit\npermission to use them. See this answer for an explanation or wikipedia or something you trust more.\nsudo setcap 'cap_net_bind_service=+ep' /opt/yourGoBinary"`
	DevMode  bool             `toml:"dev-mode"  comment:"Active develop mode"`
	Database storage.Database `toml:"database"`
}

// DefaultConfig is ...
func DefaultConfig() *Config {
	return &Config{
		DevMode:  false,
		HTTPAddr: "0.0.0.0:8088",
		Database: storage.Database{
			Storage: storage.Sqlite,
			Sqlite: sqlite.Config{
				DBPath: "./lime_base/data.db",
			},
			Postgres: postgres.Config{
				User:            "user",
				Password:        "password",
				Host:            "localhost:5430",
				Database:        "lime",
				MaxConn:         50,
				MaxIdleConn:     10,
				MaxLifetimeConn: 300,
			},
		},
	}
}

// LoadConfig is ...
func LoadConfig() (*Config, error) {
	config := DefaultConfig()

	if fsutil.IsFile(configFile) {
		file, err := os.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		if err := toml.Unmarshal(file, &config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

// SaveConfig is ...
func SaveConfig(config *Config) error {
	byteConfig, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	if err := os.WriteFile(configFile, byteConfig, 0666); err != nil {
		return err
	}

	return nil
}
