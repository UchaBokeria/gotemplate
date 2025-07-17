package general

import (
	"main/web/admin/controllers/settings/general/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	router.GET("/general", controller.Set[dtos.ReadGeneralSettingsDto](Read))
	router.PUT("/general", controller.Set[dtos.UpdateGeneralSettingsDto](Update))
}
