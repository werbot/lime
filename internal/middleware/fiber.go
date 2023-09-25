package middleware

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/rs/zerolog"
)

// FiberMiddleware is ...
func FiberMiddleware(a *fiber.App, log zerolog.Logger) {
	a.Use(
		cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
			AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With",
			ExposeHeaders:    "Origin",
			AllowCredentials: true,
		}),
		etag.New(),
	)
	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log,
	}))

}
