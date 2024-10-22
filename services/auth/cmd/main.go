package main

import (
	"fmt"
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/RafatMeraz/ecom-micro/auth/database"
	"github.com/RafatMeraz/ecom-micro/auth/internal/handler"
	"github.com/RafatMeraz/ecom-micro/auth/internal/middleware"
	"github.com/RafatMeraz/ecom-micro/auth/internal/routes"
	"github.com/RafatMeraz/ecom-micro/auth/server/http"
	"github.com/RafatMeraz/ecom-micro/pkg/logger"
	capi "github.com/hashicorp/consul/api"
	"go.uber.org/fx"
	"log"
	"log/slog"
	"os"
	"strconv"
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
		fx.Invoke(serviceRegistryWithConsul),
		fx.Invoke(middleware.SetupMiddlewares),
		fx.Invoke(routes.SetUpRoutes),
		fx.Invoke(http.StartServer),
	)
	app.Run()
}

func serviceRegistryWithConsul(cnf *configs.Config) {
	config := capi.DefaultConfig()
	consul, err := capi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	serviceID := "ecom-auth"

	port, _ := strconv.Atoi(cnf.HTTPServer.Port)
	address := getHostname()

	registration := &capi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "Ecommerce Auth Service",
		Port:    port,
		Address: address,
		Check: &capi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/health", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s", cnf.HTTPServer.ListenAddr)
	}
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}
