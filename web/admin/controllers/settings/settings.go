package settings

import (
	"main/web/admin/controllers/settings/general"
	"main/web/admin/controllers/settings/translation"

	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	settingsRouter := router.Group("/settings")

	// Register sub-components
	general.New(settingsRouter)
	translation.New(settingsRouter)
}
