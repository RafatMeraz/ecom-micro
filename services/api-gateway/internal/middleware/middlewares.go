package middleware

import (
	"github.com/RafatMeraz/ecom-micro/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupMiddlewares(app *fiber.App, loggerMiddleware *middleware.LoggerMiddleware) {
	app.Use(requestid.New())
	app.Use(loggerMiddleware.InfoLogger())
}
