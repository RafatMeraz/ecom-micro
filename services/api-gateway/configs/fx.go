package configs

import "go.uber.org/fx"

var Module = fx.Module("configs", fx.Provide(NewConfig))
