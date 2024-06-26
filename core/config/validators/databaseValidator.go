package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/himdhiman/unified-framework-go/core/config/models"
)

const (
	NoSQL = "NoSQL"
	SQL   = "SQL"
)

func isSupportedDataBase(databaseType string) bool {
	switch databaseType {
	case NoSQL, SQL:
		return true
	}
	return false
}

var ValidateDataBaseConfig validator.Func = func(fl validator.FieldLevel) bool {
	if databaseConfig, ok := fl.Field().Interface().([]models.Database); ok {
		if len(databaseConfig) == 0 {
			return true
		}
		for _, config := range databaseConfig {
			if !isSupportedDataBase(config.Type) {
				fmt.Printf("Invalid DataBase Type: %s\n", config.Type)
				return false
			}
		}
	} else {
		fmt.Printf("Error: field is not a slice of DataBaseConfig, error occured while loading DataBase Configuration")
	}
	return true
}
