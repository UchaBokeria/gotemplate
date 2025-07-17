package schema

import (
	"main/internal/models"
)

var Models = []interface{}{
	&models.Users{},
	&models.GeneralSettings{},
	&models.Translation{},
	&models.Product{},
	&models.Invoice{},
	&models.Order{},
}
