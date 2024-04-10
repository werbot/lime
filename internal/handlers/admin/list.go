package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/werbot/lime/internal/queries"
	"github.com/werbot/lime/pkg/logging"
	"github.com/werbot/lime/pkg/webutil"
)

// ListPatterns is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/list/patterns [get]
func ListPatterns(c *fiber.Ctx) error {
	log := logging.New()

	patterns, err := queries.DB().ListPatterns(c.Context())
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Patterns", patterns)
}

// ListCustomers is a ...
// @accept application/json
// @Produce application/json
// @Param
// @Success 200 {string} string "{"status":"200", "msg":""}"
// @Router /_/api/list/customers [get]
func ListCustomers(c *fiber.Ctx) error {
	log := logging.New()

	customers, err := queries.DB().ListCustomers(c.Context())
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c, nil)
	}

	return webutil.StatusOK(c, "Customers", customers)
}
