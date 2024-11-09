package configs

import (
	"github.com/RafatMeraz/ecom-micro/pkg/models"
	"github.com/RafatMeraz/ecom-micro/pkg/service"
)

type Config struct {
	HTTPServer      HTTPServerConfig
	DB              DBConfig
	RateLimitConfig models.RateLimitConfig
	PasswordHash    Hash
	TokenConfig     *service.JwtServiceConfig
}

type HTTPServerConfig struct {
	ListenAddr string
	Address    string
	Port       string
}

type DBConfig struct {
	Host     string
	Port     int
	Password string
	Name     string
	User     string
}

type Hash struct {
	Salt string
}
