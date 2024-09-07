package routes

import (
	"github.com/RafatMeraz/ecom-micro/auth/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, healthHandler *handler.HealthHandler) *fiber.App {
	router := app.Group("")

	router.Get("/health", healthHandler.GetHealth)

	return app
}
