package main

import (
	"api-gateway/configs"
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"
	"api-gateway/server/http"
	"github.com/RafatMeraz/ecom-micro/pkg/logger"
	"go.uber.org/fx"
	"log/slog"
)

func main() {
	app := fx.New(
		configs.Module,
		http.Module,
		handler.Module,
		middleware.Module,
		fx.Invoke(func() {
			authLogger := logger.NewLogger("api-gateway")
			slog.SetDefault(authLogger)
		}),
		fx.Invoke(middleware.SetupMiddlewares),
		fx.Invoke(routes.SetUpRoutes),
		fx.Invoke(http.StartServer),
	)

	app.Run()
}
