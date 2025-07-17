package products

import (
	"main/web/app/controllers/products/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	router := web.Group("/products")
	router.GET("/list", controller.Set[dtos.ProductFilterDto](ProductsList))
	router.GET("/:id", controller.Set[any](ProductDetailApp))
	router.GET("/categories", controller.Set[any](ProductsCategories))
}
