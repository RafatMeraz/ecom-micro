package middleware

import (
	"github.com/RafatMeraz/ecom-micro/pkg/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module("middleware",
	fx.Provide(middleware.NewLoggerMiddleware),
	//fx.Provide(middleware.NewRateLimiterMiddleware),
)
