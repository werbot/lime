package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

func Ping(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Pong", nil)
}
