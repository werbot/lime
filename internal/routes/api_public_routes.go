package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/manager"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	c.Get("/ping", handlers.Ping)

	api := c.Group("/api")

	// public section
	api.Get("/license/download/:id", handlers.DownloadLicense)
	api.Post("/license/verify", handlers.LicenseVerify)
}
