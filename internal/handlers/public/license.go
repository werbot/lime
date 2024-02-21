package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

// LicenseVerify is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/license/verify [post]
func LicenseVerify(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "License Verify", nil)
}

// DownloadLicense is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /api/license/download/:id [get]
func DownloadLicense(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Download License", nil)
}
