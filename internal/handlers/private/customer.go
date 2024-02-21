package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

// Customers is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/subscription/:id/*action [get]]
func Customers(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Customers", nil)
}
