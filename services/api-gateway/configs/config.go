package configs

import "github.com/RafatMeraz/ecom-micro/pkg/models"

type Config struct {
	HTTPServer      HTTPServerConfig
	RateLimitConfig models.RateLimitConfig
}

type HTTPServerConfig struct {
	ListenAddr string
}
