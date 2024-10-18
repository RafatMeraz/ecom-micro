package routes

import (
	"api-gateway/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, healthHandler *handler.HealthHandler) *fiber.App {
	router := app.Group("")

	router.Get("/health", healthHandler.GetHealth)

	return app
}
