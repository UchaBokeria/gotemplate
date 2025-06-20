package posts

import (
	"main/internal/models"
	"main/internal/storage"
)

var Type = []models.Posts_types{
	{
		Name: "ფოტო",
		Slug: "photo",
	},
	{
		Name: "ვიდეო",
		Slug: "video",
	},
	{
		Name: "ბლოგი",
		Slug: "blog",
	},
	{
		Name: "ავტოსპორტი",
		Slug: "autosport",
	},
	{
		Name: "ივენთი",
		Slug: "event",
	},
}

func Types() {
	for _, row := range Type {
		storage.DB.Create(&row)
	}
}
