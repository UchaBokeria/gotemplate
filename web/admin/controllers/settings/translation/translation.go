package translation

import (
	"main/internal/types/dto"
	"main/web/admin/controllers/settings/translation/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	translationGroup := router.Group("/translation")

	// CRUD operations
	translationGroup.GET("/list", controller.Set[dtos.TranslationListDto](list))
	translationGroup.GET("/:id", controller.Set[dto.ByID](show))
	translationGroup.POST("", controller.Set[dtos.CreateTranslationDto](create))
	translationGroup.PUT("/:id", controller.Set[dtos.UpdateTranslationDto](update))
	translationGroup.DELETE("/:id", controller.Set[dto.ByID](deleteTranslation))
}
