package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/admin"
	"github.com/werbot/lime/internal/middleware"
)

// ApiPrivateRoutes is ...
func ApiPrivateRoutes(c *fiber.App) {
	api := c.Group("/_/api")

	sign := api.Group("/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected("admin"), handlers.SignOut)

	license := api.Group("/license", middleware.JWTProtected("admin"))
	license.Get("/", handlers.Licenses)
	license.Post("/", handlers.NewLicense)
	license.Get(`/:id<regex((lic_|cst_)\w{15})>`, handlers.License)
	license.Patch("/:id", handlers.UpdateLicense)

	pattern := api.Group("/pattern", middleware.JWTProtected("admin"))
	pattern.Get("/", handlers.Patterns)
	pattern.Post("/", handlers.NewPattern)
	pattern.Get(`/:id<regex(\w{15})>`, handlers.Pattern)
	pattern.Patch("/:id", handlers.UpdatePattern)

	customer := api.Group("/customer", middleware.JWTProtected("admin"))
	customer.Get("/", handlers.Customers)
	customer.Post("/", handlers.NewCustomer)
	customer.Get(`/:id<regex(\w{15})>`, handlers.Customer)
	customer.Patch("/:id", handlers.UpdateCustomer)

	payment := api.Group("/payment", middleware.JWTProtected("admin"))
	payment.Get("/", handlers.Payments)
	payment.Get(`/:id<regex(\w{15})>`, handlers.Payment)

	audit := api.Group("/audit", middleware.JWTProtected("admin"))
	audit.Get("/", handlers.Audits)
	audit.Get(`/:id<regex(\w{15})>`, handlers.Audit)
}
