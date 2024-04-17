package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/himdhiman/unified-framework-go/core/config/models"
)

const (
	GET  = "GET"
	POST = "POST"
)

func isSupportedHandlerMethod(httpMethodType string) bool {
	switch httpMethodType {
	case GET, POST:
		return true
	}
	return false
}

var ValidateHandlerConfig validator.Func = func(fl validator.FieldLevel) bool {
	if handlerConfig, ok := fl.Field().Interface().([]models.HandlerConfig); ok {
		for _, config := range handlerConfig {
			if !isSupportedHandlerMethod(config.HTTPMethod) {
				fmt.Printf("Invalid HTTP method for handler %s: %s\n", config.Handler, config.HTTPMethod)
				return false
			}
		}
	}
	fmt.Printf("Error: field is not a slice of HandlerConfig, error occured while loading Handler Configuearion")
	return true
}

