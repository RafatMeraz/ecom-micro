package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupMiddlewares(app *fiber.App, rateLimiter *RateLimiterMiddleware, loggerMiddleware *LoggerMiddleware) {
	app.Use(requestid.New())
	app.Use(loggerMiddleware.InfoLogger())
	app.Use(rateLimiter.FixedWindow())
}
