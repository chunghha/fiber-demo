package utils

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			zap.L().Error("Oops... Server is not shutting down! Reason: %v", zap.Error(err))
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		zap.L().Error("Oops... Server is not running! Reason: %v", zap.Error(err))
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		zap.L().Error("Oops... Server is not running! Reason: %v", zap.Error(err))
	}
}
