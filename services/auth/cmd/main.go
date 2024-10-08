package main

import (
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/RafatMeraz/ecom-micro/auth/database"
	"github.com/RafatMeraz/ecom-micro/auth/internal/handler"
	"github.com/RafatMeraz/ecom-micro/auth/internal/middleware"
	"github.com/RafatMeraz/ecom-micro/auth/internal/routes"
	"github.com/RafatMeraz/ecom-micro/auth/server/http"
	"github.com/RafatMeraz/ecom-micro/pkg/logger"
	"go.uber.org/fx"
	"log/slog"
)

func main() {
	app := fx.New(
		configs.Module,
		database.Module,
		http.Module,
		handler.Module,
		middleware.Module,
		fx.Invoke(func() {
			authLogger := logger.NewLogger("auth")
			slog.SetDefault(authLogger)
		}),
		fx.Invoke(middleware.SetupMiddlewares),
		fx.Invoke(routes.SetUpRoutes),
		fx.Invoke(http.StartServer),
	)
	app.Run()
}
