package repository

import "go.uber.org/fx"

var Module = fx.Module("repository",
	fx.Provide(NewAddressRepository),
	fx.Provide(NewUserRepository))
