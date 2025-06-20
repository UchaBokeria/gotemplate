package payments

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.PaymentStatus{
	{
		Name: "აქტიური",
	},
	{
		Name: "დასრულებული",
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
