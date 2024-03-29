package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/werbot/lime/pkg/webutil"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		path := c.Path()
		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/_/api") {
			return webutil.StatusNotFound(c, utils.StatusMessage(fiber.StatusNotFound))
		}
		return c.Next()
	})
}
