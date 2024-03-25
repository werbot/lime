package webutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/internal/config"
	"github.com/werbot/lime/pkg/logging"
)

// MetaInfo is ...
type MetaInfo struct {
	Request Request `json:"request"`
	Data    any     `json:"data"`
}

// Request is ...
type Request struct {
	UserAgent   string `json:"user_agent"`
	UserIP      string `json:"user_ip"`
	UserCountry string `json:"user_country"`
}

// GetRequestInfo is ...
func GetRequestInfo(c *fiber.Ctx, data any) *MetaInfo {
	log := logging.New()
	cfg := config.Data()
	userCountry := "-"

	geoCheck := cfg.GeoDatabase.Check()
	if !geoCheck {
		country, err := cfg.GeoDatabase.GetCountryCode(c.IP())
		if err != nil {
			log.Err(err).Send()
		}
		if country.ISOCode != "" {
			userCountry = country.FlagEmoji()
		}
	}

	return &MetaInfo{
		Request: Request{
			UserAgent:   string(c.Request().Header.UserAgent()),
			UserIP:      c.IP(),
			UserCountry: userCountry,
		},
		Data: data,
	}
}
