package language

import (
	"main/internal/models"
	"main/internal/storage"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

var Translate *Placeholders
var Dictionary = make(map[string]string)

func Initialize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rawLanguage := models.Languages{}
			LangParam := parseAcceptLanguage(c.Request().Header.Get("Accept-Language"))

			storage.DB.Where(&models.Languages{Slug: LangParam}).
				Preload("LanguageDictionary").
				Last(&rawLanguage)

			// Initialize the Translate variable
			Translate = &Placeholders{}
			placeholdersValue := reflect.ValueOf(Translate).Elem()
			placeholdersType := placeholdersValue.Type()

			for _, dict := range rawLanguage.LanguageDictionary {
				Dictionary[dict.Key] = dict.Value
			}

			// Use reflection to set the fields dynamically

			for i := 0; i < placeholdersType.NumField(); i++ {
				field := placeholdersType.Field(i)
				fieldName := field.Name
				if value, found := Dictionary[fieldName]; found {
					placeholdersValue.FieldByName(fieldName).SetString(value)
				} else {
					placeholdersValue.FieldByName(fieldName).SetString("X")
				}
			}

			return next(c)
		}
	}
}

func parseAcceptLanguage(header string) string {
	languages := strings.Split(header, ",")
	for _, lang := range languages {
		lang = strings.TrimSpace(strings.Split(lang, ";")[0])
		if lang != "" {
			return lang
		}
	}
	return "en" // Fallback if nothing is found
}
