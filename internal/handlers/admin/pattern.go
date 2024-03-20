package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/errors"
	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// Patterns is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern [get]
func Patterns(c *fiber.Ctx) error {
	log := logging.New()

	pagination := webutil.GetPaginationFromCtx(c)
	patterns, err := queries.DB().Patterns(c.Context(), pagination)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Patterns", patterns)
}

// Pattern is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern/:id [get]
func Pattern(c *fiber.Ctx) error {
	log := logging.New()

	pattern, err := queries.DB().Pattern(c.Context(), c.Params("id"))
	if err != nil {
		if err == errors.ErrPatternNotFound {
			return webutil.StatusNotFound(c, errors.MsgPatternNotFound)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}
	return webutil.StatusOK(c, "Customer info", pattern)
}

// NewPattern is a ...
// @Accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern [post]
func NewPattern(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Create pattern", nil)
}

// UpdatePattern is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/pattern/:id [patch]
func UpdatePattern(c *fiber.Ctx) error {
	return webutil.StatusOK(c, "Update pattern", nil)
}
