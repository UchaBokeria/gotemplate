package types

import "main/internal/models"

type WebConfig struct {
	General      models.GeneralSettings
	Translations map[string]string
	Categories   []string
	Language     string
}
