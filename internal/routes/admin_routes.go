package routes

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/werbot/lime/web"
)

// AdminRoutes is ...
func AdminRoutes(c *fiber.App) {
	embedAdmin, _ := fs.Sub(web.EmbedAdmin(), "dist")
	c.Use("/_", filesystem.New(filesystem.Config{
		Root:         http.FS(embedAdmin),
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))
}
