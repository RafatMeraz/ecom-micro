package routes

import (
	"api-gateway/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func SetUpRoutes(app *fiber.App, healthHandler *handler.HealthHandler) *fiber.App {
	router := app.Group("")

	router.Get("/health", healthHandler.GetHealth)
	router.All("/auth/*", proxy(GetServiceUrl(AuthServiceId)))

	return app
}

func proxy(serviceURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new fasthttp request
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		// Set method and body from the original request
		req.Header.SetMethod(c.Method())
		req.SetRequestURI(serviceURL + c.Params("*"))
		req.SetBody(c.Body())

		// Copy headers from the original request
		c.Request().Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})

		// Create a fasthttp response
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Send the request to the target service
		if err := fasthttp.Do(req, resp); err != nil {
			return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		}

		// Copy the response headers and status code back to the original context
		c.Status(resp.StatusCode())
		resp.Header.CopyTo(&c.Response().Header)

		// Write the response body back to the original context
		return c.Send(resp.Body())
	}
}
