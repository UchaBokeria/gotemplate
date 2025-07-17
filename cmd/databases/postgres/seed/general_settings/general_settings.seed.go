package general_settings

import (
	"main/internal/models"
	"main/internal/storage"
)

var Seed = []models.GeneralSettings{
	{
		CompanyName: "Alfashop",
		Address:     "1234 Auto Parts Blvd, Automotive District, AD 12345",
		Phone:       "(555) 123-AUTO (2886)",
		Email:       "support@alfashop.com",
		Facebook:    "https://facebook.com/alfashop",
		Instagram:   "https://instagram.com/alfashop",
		TikTok:      "https://tiktok.com/@alfashop",
		AboutUs:     "Exclusive importer and distributor of premium automotive parts since 1995. We specialize in high-quality filters, brake components, and engine parts from world-leading manufacturers.",
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
