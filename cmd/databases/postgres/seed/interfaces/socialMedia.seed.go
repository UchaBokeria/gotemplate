package interfaces

import (
	"main/internal/models"
	"main/internal/storage"
)

var SocialMedias = []models.Social_media{
	{
		InterfaceID: 1,
		Name:        "Facebook",
		Slug:        "facebook",
		Url:         "https://www.facebook.com/alfashop",
		IconID:      4,
	},
	{
		InterfaceID: 1,
		Name:        "Instagram",
		Slug:        "instagram",
		Url:         "https://www.instagram.com/alfashop",
		IconID:      5,
	},
	{
		InterfaceID: 1,
		Name:        "Twitter",
		Slug:        "twitter",
		Url:         "https://www.twitter.com/alfashop",
		IconID:      6,
	},
	{
		InterfaceID: 1,
		Name:        "YouTube",
		Slug:        "youtube",
		Url:         "https://www.youtube.com/example",
		IconID:      2,
	},
}

func SocialMedia() {
	for _, row := range SocialMedias {
		storage.DB.Create(&row)
	}
}
