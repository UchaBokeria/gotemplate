package translation

import (
	"fmt"
	"net/http"

	"main/internal/models"
	"main/internal/storage"
	"main/internal/types/dto"
	"main/web/admin/controllers/settings/translation/dtos"
	"main/web/admin/view/components"

	"github.com/UchaBokeria/goyard/controller"
	"gorm.io/gorm"
)

func list(ctx *controller.Context[any], listDto *dtos.TranslationListDto) error {
	var translations []models.Translation
	var total int64

	query := storage.DB.Model(&models.Translation{})

	if listDto.Search != "" {
		query = query.Where("key ILIKE ? OR value ILIKE ?",
			"%"+listDto.Search+"%", "%"+listDto.Search+"%")
	}

	if listDto.Language != "" {
		query = query.Where("language = ?", listDto.Language)
	}

	query.Count(&total)

	if listDto.Page < 1 {
		listDto.Page = 1
	}
	if listDto.Limit < 1 {
		listDto.Limit = 20
	}

	offset := (listDto.Page - 1) * listDto.Limit
	query = query.Offset(offset).Limit(listDto.Limit)

	query = query.Order("key ASC")

	if err := query.Find(&translations).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to fetch translations: "+err.Error())
	}

	return ctx.Html(components.TranslationsTable(translations))
}

func show(ctx *controller.Context[any], idDto *dto.ByID) error {
	var translation models.Translation

	if err := storage.DB.First(&translation, idDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Translation not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to fetch translation: "+err.Error())
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("Translation: %s (%s) = %s", translation.Key, translation.Language, translation.Value))
}

func create(ctx *controller.Context[any], createDto *dtos.CreateTranslationDto) error {
	var existing models.Translation
	result := storage.DB.Where("key = ? AND language = ?", createDto.Key, createDto.Language).First(&existing)
	if result.Error == nil {
		return ctx.String(http.StatusConflict, "Translation with this key and language already exists")
	}

	translation := models.Translation{
		Key:      createDto.Key,
		Value:    createDto.Value,
		Language: createDto.Language,
	}

	if err := storage.DB.Create(&translation).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create translation: "+err.Error())
	}

	// Return updated translations list for the same language
	var translations []models.Translation
	storage.DB.Where("language = ?", createDto.Language).Order("key ASC").Find(&translations)
	return ctx.Html(components.TranslationsTable(translations))
}

func update(ctx *controller.Context[any], updateDto *dtos.UpdateTranslationDto) error {
	var translation models.Translation

	if err := storage.DB.First(&translation, updateDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Translation not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find translation: "+err.Error())
	}

	if translation.Key != updateDto.Key || translation.Language != updateDto.Language {
		var existing models.Translation
		result := storage.DB.Where("key = ? AND language = ? AND id != ?", updateDto.Key, updateDto.Language, updateDto.ID).First(&existing)
		if result.Error == nil {
			return ctx.String(http.StatusConflict, "Translation with this key and language already exists")
		}
	}

	translation.Key = updateDto.Key
	translation.Value = updateDto.Value
	translation.Language = updateDto.Language

	if err := storage.DB.Save(&translation).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update translation: "+err.Error())
	}

	// Return updated translations list for the same language
	var translations []models.Translation
	storage.DB.Where("language = ?", updateDto.Language).Order("key ASC").Find(&translations)
	return ctx.Html(components.TranslationsTable(translations))
}

func deleteTranslation(ctx *controller.Context[any], deleteDto *dto.ByID) error {
	var translation models.Translation

	if err := storage.DB.First(&translation, deleteDto.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.String(http.StatusNotFound, "Translation not found")
		}
		return ctx.String(http.StatusInternalServerError, "Failed to find translation: "+err.Error())
	}

	language := translation.Language

	if err := storage.DB.Delete(&translation).Error; err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete translation: "+err.Error())
	}

	// Return updated translations list for the same language
	var translations []models.Translation
	storage.DB.Where("language = ?", language).Order("key ASC").Find(&translations)
	return ctx.Html(components.TranslationsTable(translations))
}
