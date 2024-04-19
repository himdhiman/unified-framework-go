package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/himdhiman/unified-framework-go/core/config/models"
	"github.com/himdhiman/unified-framework-go/core/config/validators"
	"github.com/spf13/viper"
)

func LoadConfig(filename string, customValidators map[string]validator.Func) (*models.Config, error) {
	viper.SetConfigFile(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg models.Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling configureation: %w", err)
	}

	validate := validator.New()

	if customValidators == nil {
		customValidators = make(map[string]validator.Func)
	}

	customValidators["handlerconfig"] = validators.ValidateHandlerConfig
	customValidators["databaseconfig"] = validators.ValidateDataBaseConfig

	for tag, customValidator := range customValidators {
		if err := validate.RegisterValidation(tag, customValidator); err != nil {
			return nil, fmt.Errorf("error registering custom validator '%s': %w", tag, err)
		}
	}

	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return &cfg, nil
}
