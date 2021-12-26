package controllers

import (
	"fiber-demo/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetNewAccessToken method to respond a new access token.
// @Description Respond a new access token.
// @Summary respond a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 on error of token generation.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
