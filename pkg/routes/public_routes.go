package routes

import (
	"fiber-demo/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func to serve group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
