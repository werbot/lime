package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

// GetKey is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /key/:customer_id [get]
func GetKey(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "GetKey", nil)
}

// CreateKey is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /key [post]
func CreateKey(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "CreateKey", nil)
}

// UpdateKey is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /key/:customer_id [PATCH]
func UpdateKey(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "UpdateKey", nil)
}

// VerifyKey is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /verify [post]
func VerifyKey(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "VerifyKey", nil)
}
