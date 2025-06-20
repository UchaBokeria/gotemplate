package interfaces

import (
	"main/internal/models"
	"main/internal/storage"
)

var Languages = []models.Languages{
	{
		InterfaceID: 1,
		IconID:      45,
		Name:        "ქართული",
		Slug:        "ge",
	},
	{
		InterfaceID: 1,
		IconID:      46,
		Name:        "English",
		Slug:        "en",
	},
	{
		InterfaceID: 1,
		IconID:      47,
		Name:        "Русский",
		Slug:        "ru",
	},
}

var LanguageDictionary = []models.LanguageDictionary{
	{
		LanguageID: 4,
		Key:        "NewsPageTitle",
		Value:      "სიახლეები",
	},
	{
		LanguageID: 5,
		Key:        "NewsPageTitle",
		Value:      "News",
	},
	{
		LanguageID: 6,
		Key:        "NewsPageTitle",
		Value:      "სიახლეები",
	},
}

func Language() {
	for _, lang := range Languages {
		storage.DB.Create(&lang)
	}
	for _, dict := range LanguageDictionary {
		storage.DB.Create(&dict)
	}
}
