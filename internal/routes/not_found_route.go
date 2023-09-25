package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/pkg/webutil"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return webutil.StatusNotFound(c)
		}
		return c.Next()
	})
}
