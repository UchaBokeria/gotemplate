package categories

import (
	"main/internal/types/dto"
	"main/web/admin/controllers/categories/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	categoriesGroup := router.Group("/categories")

	// CRUD operations
	categoriesGroup.GET("/list", controller.Set[dtos.CategoryListDto](list))
	categoriesGroup.GET("/tree", controller.Set[dtos.CategoryListDto](tree))
	categoriesGroup.GET("/:id", controller.Set[dto.ByID](show))
	categoriesGroup.POST("", controller.Set[dtos.CreateCategoryDto](create))
	categoriesGroup.PUT("/:id", controller.Set[dtos.UpdateCategoryDto](update))
	categoriesGroup.DELETE("/:id", controller.Set[dto.ByID](deleteCategory))
	
	// Additional tree operations
	categoriesGroup.GET("/:id/children", controller.Set[dto.ByID](getChildren))
	categoriesGroup.POST("/:id/move", controller.Set[dtos.UpdateCategoryDto](moveCategory))
} 