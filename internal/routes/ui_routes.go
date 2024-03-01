package routes

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/werbot/lime/web"
)

// UIRoutes is ...
func UIRoutes(c *fiber.App) {
	embedAdmin, _ := fs.Sub(web.EmbedUI(), "dist")
	c.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(embedAdmin),
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))
}
