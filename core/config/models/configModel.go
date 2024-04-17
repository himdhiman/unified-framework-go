package models

type Config struct {
	ApplicationName string          `mapstructure:"applicationName" validate:"required"`
	ServiceName     string          `mapstructure:"serviceName" validate:"required"`
	Environment     string          `mapstructure:"environemnt" validate:"required"`
	Version         float32         `mapstructure:"version" validate:"required"`
	Port            int             `mapstructure:"port" validate:"required"`
	Handlers        []HandlerConfig `mapstructure:"handlers"`
	Database        []Database      `mapstructure:"database"`
}
