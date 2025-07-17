package middlewares

import (
	"main/internal/models"
	"main/internal/storage"
	"main/web/app/types"
	"strings"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func WebConfig() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Use[any](func(ctx *controller.Context[any]) error {
			// Get language from Accept-Language header or default to "en"
			language := ctx.Request().Header.Get("Accept-Language")
			if strings.Contains(language, "ka") {
				language = "ka"
			} else if strings.Contains(language, "ru") {
				language = "ru"
			} else {
				language = "en"
			}
			// Fetch translations from database
			var translations []models.Translation
			err := storage.DB.Where("language = ?", language).Find(&translations).Error
			if err != nil {
				// If error, continue without translations
				return next(ctx)
			}

			// Convert to key-value map for easier use
			translationMap := make(map[string]string)
			for _, t := range translations {
				translationMap[t.Key] = t.Value
			}

			var generalSettings models.GeneralSettings
			err = storage.DB.Last(&generalSettings).Error
			if err != nil {
				return next(ctx)
			}

			// var categories []models.Category
			// err = storage.DB.Find(&categories).Error
			// if err != nil {
			// 	return next(ctx)
			// }
			var categories []string = []string{"Auto Parts", "Auto Accessories", "Auto Electronics", "Auto Tools", "Auto Parts"}

			var WEBCONFIG = types.WebConfig{
				General:      generalSettings,
				Translations: translationMap,
				Categories:   categories,
				Language:     language,
			}
			// Set translations in context for handlers to use
			ctx.Set("webconfig", WEBCONFIG)

			return next(ctx)
		})
	}
}
