package main

import (
	"fiber-demo/pkg/configs"
	"fiber-demo/pkg/middleware"
	"fiber-demo/pkg/routes"
	"fiber-demo/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	_ "fiber-demo/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var logger *zap.Logger

func initLogger() {
	// global zap logger
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

// @title fiber-demo
// @version 0.1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	initLogger()

	config := configs.FiberConfig()

	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.HealthRoute(app)   // Register a route for Health report.
	routes.SwaggerRoute(app)  // Register a route for Swagger
	routes.PublicRoutes(app)  // Register a public routes for app
	routes.NotFoundRoute(app) // Place this at last, register a route for 404 error

	utils.StartServer(app)
}
