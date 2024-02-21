package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

// ManageLicense is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/sign/in [post]]
func ManageLicense(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Manage License", nil)
}

// AccessLink is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/sign/out [get]
func AccessLink(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Access Link", nil)
}
