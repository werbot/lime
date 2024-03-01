package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/public"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	c.Get("/ping", handlers.Ping)

	api := c.Group("/api")

	sign := api.Group("/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out" /*middleware.JWTProtected(),*/, handlers.SignOut)

	manager := api.Group("/manager")
	manager.Post("/license", handlers.ManageLicense)

	license := api.Group("/license")
	license.Get("/download/:id", handlers.DownloadLicense)
	license.Post("/verify", handlers.LicenseVerify)
}
