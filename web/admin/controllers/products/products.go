package products

import (
	"main/internal/types/dto"
	"main/web/admin/controllers/products/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	productsGroup := router.Group("/products")

	// CRUD operations
	productsGroup.GET("/list", controller.Set[dtos.ProductListDto](list))
	productsGroup.GET("/:id", controller.Set[dto.ByID](show))
	productsGroup.POST("", controller.Set[dtos.CreateProductDto](create))
	productsGroup.PUT("/:id", controller.Set[dtos.UpdateProductDto](update))
	productsGroup.DELETE("/:id", controller.Set[dto.ByID](deleteProduct))
}
