package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/manager"
	"github.com/werbot/lime/internal/middleware"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	c.Get("/ping", handlers.Ping)

	api := c.Group("/api")

	// public section
	api.Get("/license/download/:id", handlers.DownloadLicense)
	api.Post("/license/verify", handlers.LicenseVerify)

	// manager section
	sign := api.Group("/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected("manager"), handlers.SignOut)

	manager := api.Group("/manager", middleware.JWTProtected("manager"))
	manager.Post("/license", handlers.ManageLicense)

	license := api.Group("/license", middleware.JWTProtected("manager"))
	license.Get("/", handlers.Licenses)
	license.Get(`/:id<regex((lic_)\w{15})>`, handlers.License)
}
