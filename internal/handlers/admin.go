package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/werbot/lime/pkg/webutil"
)

func SignIn(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "SignIn", nil)
}

func SignOut(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "SignOut", nil)
}

func MainHandler(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "MainHandler", nil)
}

func CustomerList(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "CustomerList", nil)
}
