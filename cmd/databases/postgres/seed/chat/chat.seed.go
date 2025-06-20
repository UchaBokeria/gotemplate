package chat

import (
	"main/internal/models"
	"main/internal/storage"
)

var ChatStatus []models.Chat_status = []models.Chat_status{
	{
		Name: "აქტიური",
		Slug: "open",
	},
	{
		Name: "დასრულებული",
		Slug: "closed",
	},
}

var ChatType []models.Chat_status = []models.Chat_status{
	{
		Name: "ლაივი",
		Slug: "live",
	},
	{
		Name: "ელფოსტა",
		Slug: "mail",
	},
}

func Populate() {
	for _, row := range ChatType {
		storage.DB.Create(&row)
	}
	for _, row := range ChatStatus {
		storage.DB.Create(&row)
	}
}
