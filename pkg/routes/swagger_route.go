package routes

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// SwaggerRoute func to serve swagger API doc route.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	route.Get("*", swagger.HandlerDefault)
}
