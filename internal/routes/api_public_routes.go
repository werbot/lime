package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/handlers"
	handlersV1 "github.com/werbot/lime/internal/handlers/v1"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	c.Get("/ping", handlers.Ping)

	apiV1 := c.Group("/api/v1")
	apiV1.Post("/key", handlersV1.CreateKey)       // controllers.CreateKey
	apiV1.Get("/key/:key", handlersV1.GetKey)      // controllers.GetKey
	apiV1.Patch("/key/:key", handlersV1.UpdateKey) // controllers.UpdateKey
	apiV1.Post("/verify", handlersV1.VerifyKey)    // controllers.VerifyKey
}
