package configs

import (
	"fmt"
	"github.com/RafatMeraz/ecom-micro/pkg/models"
	"github.com/RafatMeraz/ecom-micro/pkg/service"
	"github.com/spf13/viper"
)

func getViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.SetConfigName("config")
	return v
}

func NewConfig() (*Config, error) {
	fmt.Println("loading config")
	v := getViper()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	config = Config{
		HTTPServer: HTTPServerConfig{
			ListenAddr: v.GetString("HttpServer.listenAddr"),
			Address:    v.GetString("HttpServer.address"),
			Port:       v.GetString("HttpServer.port"),
		},
		DB: DBConfig{
			Host:     v.GetString("DB.host"),
			Port:     v.GetInt("DB.port"),
			User:     v.GetString("DB.user"),
			Password: v.GetString("DB.password"),
			Name:     v.GetString("DB.name"),
		},
		RateLimitConfig: models.RateLimitConfig{
			RequestsPerMin: v.GetInt("RateLimit.limitPerMin"),
		},
		PasswordHash: Hash{
			Salt: v.GetString("Hash.passwordSalt"),
		},
		TokenConfig: &service.JwtServiceConfig{
			JwtSecret:               v.GetString("Jwt.secretKey"),
			RefreshJwtSecret:        v.GetString("Jwt.refreshTokenSecretKey"),
			TokenExpireAtDay:        v.GetInt("Jwt.tokenExpireDay"),
			RefreshTokenExpireAtDay: v.GetInt("Jwt.refreshTokenExpireDay"),
		},
	}

	return &config, nil
}
