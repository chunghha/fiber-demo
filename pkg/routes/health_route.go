package routes

import (
	"fiber-demo/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// HealthRoute func to serve group of health routes.
func HealthRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")

	route.Get("/is_healthy", controllers.IsHealthy) // report api up
}
