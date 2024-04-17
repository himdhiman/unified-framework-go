package models

type Database struct {
	Name     string `mapstructure:"name" validate:"required"`
	Type     string `mapstructure:"type" validate:"required,databasetypeconfig"`
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
}
