package schema

import (
	"main/internal/models"
)

var Models = []interface{}{
	&models.Category{},
	&models.Users{},
	&models.GeneralSettings{},
	&models.Translation{},
	&models.Product{},
	&models.ProductCategory{},
	&models.Invoice{},
	&models.Order{},
}
