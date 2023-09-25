package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/internal/handlers"
	"github.com/werbot/lime/internal/middleware"
)

// ApiPrivateRoutes is ...
func ApiPrivateRoutes(c *fiber.App) {
	root := c.Group("/_")

	sign := root.Group("/api/sign")
	sign.Post("/in", handlers.SignIn)                              // middleware.Login
	sign.Post("/out", middleware.JWTProtected(), handlers.SignOut) // middleware.Logout

	api := root.Group("/api", middleware.JWTProtected())
	api.Get("/", handlers.MainHandler)                          // controllers.MainHandler
	api.Get("/subscription/:id/*action", handlers.CustomerList) // controllers.CustomerSubscrptionList
	//api.Get("/license/:id", handlers.DownloadLicense)                      // controllers.DownloadLicense
}
