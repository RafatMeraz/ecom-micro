package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"net/http"
	"time"
)

type RateLimiterMiddleware struct {
	limitPerMinute int
}

func NewRateLimiterMiddleware(limitPerMin int) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		limitPerMinute: limitPerMin,
	}
}

func (r RateLimiterMiddleware) FixedWindow() fiber.Handler {
	return limiter.New(
		limiter.Config{
			Max:               r.limitPerMinute,
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
			Max:               r.limitPerMinute,
			Expiration:        1 * time.Minute,
			LimiterMiddleware: limiter.SlidingWindow{},
			LimitReached: func(ctx *fiber.Ctx) error {
				return ctx.SendStatus(http.StatusTooManyRequests)
			},
		},
	)
}
