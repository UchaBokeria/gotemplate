package middlewares

import (
	"main/internal/models"
	"main/internal/storage"
	"main/web/app/types"
	"main/web/app/view"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func WebConfig() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			language := "en"
			cookie, err := ctx.Cookie("lang")
			if err == nil && cookie != nil {
				language = cookie.Value
			}

			var translations []models.Translation
			if err := storage.DB.Where("language = ?", language).Find(&translations).Error; err != nil {
				return next(ctx)
			}

			// Convert to key-value map for easier use
			translationMap := make(map[string]string)
			for _, t := range translations {
				translationMap[t.Key] = t.Value
			}

			var generalSettings models.GeneralSettings
			if err := storage.DB.Last(&generalSettings).Error; err != nil {
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
			ctx.Set("LayoutRenderNoHtmx", func(childComponent templ.Component) templ.Component {
				parts := strings.Split(ctx.Request().URL.Path, "/")
				if len(parts) > 1 && parts[1] == "admin" {
					return childComponent
				}
				return view.Index(WEBCONFIG, childComponent)
			})
			ctx.Set("webconfig", WEBCONFIG)
			return next(ctx)
		}
	}
}
