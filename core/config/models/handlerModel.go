package models

type HandlerConfig struct {
	Endpoint   string `mapstructure:"endpoint" validate:"required"`
	Handler    string `mapstructure:"handler" validate:"required"`
	HTTPMethod string `mapstructure:"httpMethod" validate:"required"`
}
