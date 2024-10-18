package configs

import (
	"fmt"
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
		},
		RateLimit: RateLimit{
			LimitPerMin: v.GetInt("RateLimit.limitPerMin"),
		},
	}

	return &config, nil
}
