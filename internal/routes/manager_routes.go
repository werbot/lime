package routes

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/werbot/lime/web"
)

// AdminRoutes is ...
func ManagerRoutes(c *fiber.App) {
	embedAdmin, _ := fs.Sub(web.EmbedAdmin(), "dist")
	c.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(embedAdmin),
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))
}
