package schema

import (
	"main/internal/models"
)

var Models = []interface{}{
	&models.Users{},
	&models.GeneralSettings{},
	&models.Translation{},
	&models.Category{},
	&models.Product{},
	&models.ProductCategory{},
	&models.Invoice{},
	&models.Order{},
}
