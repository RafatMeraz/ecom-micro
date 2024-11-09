package service

import (
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"github.com/RafatMeraz/ecom-micro/pkg/service"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(newJwtTokenService),
	fx.Provide(NewUserService),
)

func newJwtTokenService(config *configs.Config) *service.JwtService {
	return service.NewJwtService(config.TokenConfig)
}
