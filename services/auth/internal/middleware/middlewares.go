package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddlewares(app *fiber.App, rateLimiter *RateLimiterMiddleware) {
	app.Use(logger.New())
	app.Use(rateLimiter.FixedWindow())
}
