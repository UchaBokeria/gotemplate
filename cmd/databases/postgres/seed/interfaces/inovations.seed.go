package interfaces

import (
	"main/internal/models"
	"main/internal/storage"
)

var Inovations = []models.Interface_inovations{
	{
		InterfaceID: 1,
		Name:        "USVO",
		Slug:        "usvo",
		Title:       "USVO",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/usvo",
		PicID:       Pint(7),
	},
	{
		InterfaceID: 1,
		Name:        "CleanSynto",
		Slug:        "cleansynto",
		Title:       "Clean Synto",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/cleansynto",
		PicID:       Pint(8),
	},
	{
		InterfaceID: 1,
		Name:        "ATFprofessionalLine",
		Slug:        "atfprofessionalline",
		Title:       "ATF პროფესიონალური",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/atf-professional-line",
		PicID:       Pint(9),
	},
	{
		InterfaceID: 1,
		Name:        "BoxInBag",
		Slug:        "boxinbag",
		Title:       "ჩანთა ყუთში",
		Url:         "https://www.alfashop.de/de/unternehmen/innovationen/bag-in-box",
		PicID:       Pint(10),
	},
}

func Inovation() {
	for _, row := range Inovations {
		storage.DB.Create(&row)
	}
}
