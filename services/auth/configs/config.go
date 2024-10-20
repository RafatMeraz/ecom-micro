package configs

import "github.com/RafatMeraz/ecom-micro/pkg/models"

type Config struct {
	HTTPServer      HTTPServerConfig
	DB              DBConfig
	RateLimitConfig models.RateLimitConfig
}

type HTTPServerConfig struct {
	ListenAddr string
}

type DBConfig struct {
	Host     string
	Port     int
	Password string
	Name     string
	User     string
}
