package orders

import (
	"main/internal/types/dto"
	"main/web/admin/controllers/orders/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	ordersGroup := router.Group("/orders")

	// List and search orders
	ordersGroup.GET("/list", controller.Set[dtos.OrderListDto](list))

	// Get single order
	ordersGroup.GET("/:id", controller.Set[dto.ByID](show))

	// Delete order
	ordersGroup.DELETE("/:id", controller.Set[dto.ByID](deleteOrder))
}
