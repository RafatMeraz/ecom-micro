package http

import (
	"github.com/RafatMeraz/ecom-micro/auth/internal/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartServer() {
	host := "localhost"
	port := "8080"

	app := fiber.New()

	routes.SetUpRoutes(app)

	log.Fatal(app.Listen(host + ":" + port))
}
