package general

import (
	"net/http"

	"main/internal/models"
	"main/internal/storage"
	"main/web/admin/controllers/settings/general/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
)

func Read(ctx *controller.Context[any], dto *dtos.ReadGeneralSettingsDto) error {
	var settings models.GeneralSettings

	// Get or create general settings (there should be only one record)
	result := storage.DB.First(&settings)
	if result.Error != nil {
		// Create default settings if none exist
		settings = models.GeneralSettings{
			CompanyName: "Alfashop",
			Address:     "1234 Auto Parts Blvd, Automotive District, AD 12345",
			Phone:       "(555) 123-AUTO (2886)",
			Email:       "support@alfashop.com",
			Facebook:    "https://facebook.com/alfashop",
			Instagram:   "https://instagram.com/alfashop",
			TikTok:      "https://tiktok.com/@alfashop",
			AboutUs:     "Exclusive importer and distributor of premium automotive parts since 1995.",
		}

		if err := storage.DB.Create(&settings).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to create default settings: "+err.Error())
		}
	}

	return ctx.Html(components.GeneralSettingsForm(settings))
}

func Update(ctx *controller.Context[any], dto *dtos.UpdateGeneralSettingsDto) error {

	var settings models.GeneralSettings

	// Get existing settings or create new one
	result := storage.DB.First(&settings)
	if result.Error != nil {
		// Create new settings
		settings = models.GeneralSettings{}
	}

	// Update fields
	settings.CompanyName = dto.CompanyName
	settings.Address = dto.Address
	settings.Phone = dto.Phone
	settings.Email = dto.Email
	settings.Facebook = dto.Facebook
	settings.Instagram = dto.Instagram
	settings.TikTok = dto.TikTok
	settings.AboutUs = dto.AboutUs

	// Save or update
	if settings.ID == 0 {
		if err := storage.DB.Create(&settings).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to create settings: "+err.Error())
		}
	} else {
		if err := storage.DB.Save(&settings).Error; err != nil {
			return ctx.String(http.StatusInternalServerError, "Failed to update settings: "+err.Error())
		}
	}

	return ctx.Html(components.GeneralSettingsForm(settings))
}
