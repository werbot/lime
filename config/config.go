package config

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var (
	defaultConfig *viper.Viper
	DB            *gorm.DB
)

func Config() Provider {
	return defaultConfig
}

func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	var err error
	defaultConfig = readViperConfig("LIME")

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", defaultConfig.GetString("db_host"), defaultConfig.GetString("db_port"), defaultConfig.GetString("db_user"), defaultConfig.GetString("db_name"), defaultConfig.GetString("db_password")))
	if err != nil {
		fmt.Printf("Cannot connect to postgres database")
		log.Fatal("This is the error:", err)
	}
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	v.SetDefault("mode", "debug") // release
	v.SetDefault("port", ":8080")

	v.SetDefault("api_secret", "TGq7dTjt@G.vkuDYwQfdf7uZvmwr@MzV.r2r6NGtPF")

	v.SetDefault("db_host", "127.0.0.1")
	v.SetDefault("db_user", "postgres")
	v.SetDefault("db_password", "postgres")
	v.SetDefault("db_name", "license")
	v.SetDefault("db_port", "5432")

	return v
}
