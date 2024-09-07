package http

import (
	"context"
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"log"
)

func NewServer() *fiber.App {
	app := fiber.New()
	return app
}

func StartServer(lc fx.Lifecycle, app *fiber.App, config *configs.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting server...")
			go func() {
				if err := app.Listen(config.HTTPServer.ListenAddr); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			return app.Shutdown()
		},
	})
}
