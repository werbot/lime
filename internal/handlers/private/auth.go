package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

// SignIn is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/sign/in [post]]
func SignIn(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "SignIn", nil)
}

// SignOut is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/sign/out [post]
func SignOut(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "SignOut", nil)
}
