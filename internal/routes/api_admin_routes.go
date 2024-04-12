package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/werbot/lime/internal/handlers/admin"
	"github.com/werbot/lime/internal/middleware"
)

// ApiAdminRoutes is ...
func ApiAdminRoutes(c *fiber.App) {
	api := c.Group("/_/api")

	sign := api.Group("/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected("admin"), handlers.SignOut)

	license := api.Group("/license", middleware.JWTProtected("admin"))
	license.Get("/", handlers.Licenses)
	license.Post("/", handlers.NewLicense)
	license.Get(`/:id<regex((lic_|cst_)\w{15})>`, handlers.License)
	license.Patch(`/:id<regex(\w{15})>`, handlers.UpdateLicense)

	pattern := api.Group("/pattern", middleware.JWTProtected("admin"))
	pattern.Get("/", handlers.Patterns)
	pattern.Post("/", handlers.AddPattern)
	pattern.Post(`/:id<regex(\w{15})>`, handlers.ClonePattern)
	pattern.Get(`/:id<regex(\w{15})>`, handlers.Pattern)
	pattern.Patch(`/:id<regex(\w{15})>`, handlers.UpdatePattern)
	pattern.Delete(`/:id<regex(\w{15})>`, handlers.DeletePattern)

	customer := api.Group("/customer", middleware.JWTProtected("admin"))
	customer.Get("/", handlers.Customers)
	customer.Post("/", handlers.AddCustomer)
	customer.Get(`/:id<regex(\w{15})>`, handlers.Customer)
	customer.Patch(`/:id<regex(\w{15})>`, handlers.UpdateCustomer)
	customer.Delete(`/:id<regex(\w{15})>`, handlers.DeleteCustomer)

	payment := api.Group("/payment", middleware.JWTProtected("admin"))
	payment.Get("/", handlers.Payments)
	payment.Post("/", handlers.AddPayment)
	payment.Get(`/:id<regex(\w{15})>`, handlers.Payment)
	payment.Patch(`/:id<regex(\w{15})>`, handlers.UpdatePayment)

	audit := api.Group("/audit", middleware.JWTProtected("admin"))
	audit.Get("/", handlers.Audits)
	audit.Get(`/:id<regex(\w{15})>`, handlers.Audit)

	list := api.Group("/list", middleware.JWTProtected("admin"))
	list.Get("/patterns/:name?", handlers.ListPatterns)
	list.Get("/customers/:name?", handlers.ListCustomers)

	setting := api.Group("/setting", middleware.JWTProtected("admin"))
	setting.Get("/:group", handlers.SettingGroup)
	setting.Patch("/:group", handlers.UpdateSettingGroup)
}
