package types

import "main/internal/models"

type WebConfig struct {
	General      models.GeneralSettings
	Translations map[string]string
	Categories   []string
	Language     string
}

type WebDto struct {
	WebConfig
	Redirect string `query:"redirect"`
	Lang     string `param:"lang"`
}
