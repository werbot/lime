package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/private"
	"github.com/werbot/lime/internal/middleware"
)

// ApiPrivateRoutes is ...
func ApiPrivateRoutes(c *fiber.App) {
	api := c.Group("/_/api" /*, middleware.JWTProtected()*/)

	sign := api.Group("/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected(), handlers.SignOut)

	subscription := api.Group("/subscription")
	subscription.Get("/:id/*action", handlers.Customers)

	license := api.Group("/license")
	license.Post("/", handlers.NewLicense)
	license.Get("/:id", handlers.GetLicense)
	license.Patch("/:id", handlers.UpdateLicense)
}