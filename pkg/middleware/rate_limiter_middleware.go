package middleware

import (
	"github.com/RafatMeraz/ecom-micro/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"net/http"
	"time"
)

type RateLimiterMiddleware struct {
	rateLimitConfig models.RateLimitConfig
}

func NewRateLimiterMiddleware(rateLimitConfig models.RateLimitConfig) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		rateLimitConfig: rateLimitConfig,
	}
}

func (r RateLimiterMiddleware) FixedWindow() fiber.Handler {
	return limiter.New(
		limiter.Config{
			Max:               r.rateLimitConfig.RequestsPerMin,
			Expiration:        1 * time.Minute,
			LimiterMiddleware: limiter.FixedWindow{},
			LimitReached: func(ctx *fiber.Ctx) error {
				return ctx.SendStatus(http.StatusTooManyRequests)
			},
		},
	)
}

func (r RateLimiterMiddleware) SlidingWindow() fiber.Handler {
	return limiter.New(
		limiter.Config{
			Max:               r.rateLimitConfig.RequestsPerMin,
			Expiration:        1 * time.Minute,
			LimiterMiddleware: limiter.SlidingWindow{},
			LimitReached: func(ctx *fiber.Ctx) error {
				return ctx.SendStatus(http.StatusTooManyRequests)
			},
		},
	)
}
