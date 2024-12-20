package routes

import (
	"github.com/RafatMeraz/ecom-micro/auth/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(
	app *fiber.App,
	healthHandler *handler.HealthHandler,
	authHandler *handler.AuthHandler) *fiber.App {
	router := app.Group("")

	router.Get("/health", healthHandler.GetHealth)
	router.Post("/sign-up", authHandler.SignUp)
	router.Post("/sign-in", authHandler.SignIn)

	return app
}
