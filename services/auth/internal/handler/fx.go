package handler

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(NewValidator),
	fx.Provide(NewHealthHandler),
	fx.Provide(NewAuthHandler),
)

func NewValidator() *validator.Validate {
	return validator.New()
}
