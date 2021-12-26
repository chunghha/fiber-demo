package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// IsHealthy method to return api health report.
// @Description Respond api health report.
// @Summary respond api health report
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {boolean} status "ok"
// @Router /is_healthy [get]
func IsHealthy(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"is_healthy": true,
	})
}
