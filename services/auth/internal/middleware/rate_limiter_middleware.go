package middleware

import (
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"net/http"
	"time"
)

type RateLimiterMiddleware struct {
	config *configs.Config
}

func NewRateLimiterMiddleware(cfg *configs.Config) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		config: cfg,
	}
}

func (r RateLimiterMiddleware) FixedWindow() fiber.Handler {
	return limiter.New(
		limiter.Config{
			Max:               r.config.RateLimit.LimitPerMin,
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
			Max:               r.config.RateLimit.LimitPerMin,
			Expiration:        1 * time.Minute,
			LimiterMiddleware: limiter.SlidingWindow{},
			LimitReached: func(ctx *fiber.Ctx) error {
				return ctx.SendStatus(http.StatusTooManyRequests)
			},
		},
	)
}
