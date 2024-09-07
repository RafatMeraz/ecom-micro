package main

import (
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/RafatMeraz/ecom-micro/auth/internal/handler"
	"github.com/RafatMeraz/ecom-micro/auth/internal/routes"
	"github.com/RafatMeraz/ecom-micro/auth/server/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		configs.Module,
		http.Module,
		handler.Module,
		fx.Invoke(routes.SetUpRoutes),
		fx.Invoke(http.StartServer),
	)
	app.Run()
}
