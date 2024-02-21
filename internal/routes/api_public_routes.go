package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/public"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	c.Get("/ping", handlers.Ping)

	api := c.Group("/api")

	manager := api.Group("/manager")
	manager.Post("/license", handlers.ManageLicense)
	manager.Get("/access-link", handlers.AccessLink)

	license := api.Group("/license")
	license.Get("/download/:id", handlers.DownloadLicense)
	license.Post("/verify", handlers.LicenseVerify)
}
