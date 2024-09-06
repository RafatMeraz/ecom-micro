package main

import (
	"github.com/RafatMeraz/ecom-micro/auth/server/http"
	"go.uber.org/fx"
)

func main() {
	fx.New().Run()
	http.StartServer()
}
