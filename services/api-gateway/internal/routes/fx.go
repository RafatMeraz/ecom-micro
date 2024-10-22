package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Module("routes", fx.Invoke(SetUpRoutes))
