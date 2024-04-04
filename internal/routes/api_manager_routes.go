package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/manager"
	"github.com/werbot/lime/internal/middleware"
)

// ApiManagerRoutes is ...
func ApiManagerRoutes(c *fiber.App) {
	api := c.Group("/api")

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
