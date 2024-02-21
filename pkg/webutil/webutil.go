package webutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// HTTPResponse represents response body of API
type HTTPResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}

// Response is a takes in a Fiber context object, an HTTP status code, a message string and some data.
func Response(c *fiber.Ctx, code int, message string, data any) error {
	if message == "" {
		return c.Status(code).JSON(data)
	}

	return c.Status(code).JSON(HTTPResponse{
		Code:    code,
		Message: message,
		Result:  data,
	})
}

// StatusOK is ...
// 200 ok
func StatusOK(c *fiber.Ctx, message string, data any) error {
	return Response(c, fiber.StatusOK, message, data)
}

// StatusBadRequest is ...
// 400 error
func StatusBadRequest(c *fiber.Ctx, data any) error {
	return Response(c, fiber.StatusBadRequest, utils.StatusMessage(fiber.StatusBadRequest), data)
}

// StatusUnauthorized is ...
// 401 error
func StatusUnauthorized(c *fiber.Ctx, data any) error {
	return Response(c, fiber.StatusUnauthorized, utils.StatusMessage(fiber.StatusUnauthorized), data)
}

// StatusNotFound is ...
// 404 error
func StatusNotFound(c *fiber.Ctx, data any) error {
	return Response(c, fiber.StatusNotFound, utils.StatusMessage(fiber.StatusNotFound), data)
}

// StatusInternalServerError is ...
// 500 error
func StatusInternalServerError(c *fiber.Ctx, data any) error {
	return Response(c, fiber.StatusInternalServerError, utils.StatusMessage(fiber.StatusInternalServerError), data)
}
