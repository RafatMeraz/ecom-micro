package middleware

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"log/slog"
)

type LoggerMiddleware struct{}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) InfoLogger() fiber.Handler {
	loggerHandler := fiberLogger.New(fiberLogger.Config{
		Output: nil,
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() >= 200 && c.Response().StatusCode() < 300 {
				slog.Info("request",
					"request-id", c.GetRespHeader("X-Request-ID"),
					"url", c.OriginalURL(),
					"method", c.Method(),
					"params", c.AllParams(),
					"queryParams", c.Queries(),
					"status", c.Response().StatusCode(), // Log the response status
				)
			} else {
				slog.Error("request",
					"request-id", c.GetRespHeader("X-Request-ID"),
					"url", c.OriginalURL(),
					"method", c.Method(),
					"params", c.AllParams(),
					"queryParams", c.Queries(),
					"status", c.Response().StatusCode(),
					"error", c.Next(),
				)
			}
		},
	})
	return loggerHandler
}
