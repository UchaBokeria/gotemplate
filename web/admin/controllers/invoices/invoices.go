package invoices

import (
	"main/internal/types/dto"
	"main/web/admin/controllers/invoices/dtos"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func New(router *echo.Group) {
	invoicesGroup := router.Group("/invoices")

	// List and search invoices
	invoicesGroup.GET("/list", controller.Set[dtos.InvoiceListDto](list))

	// Get single invoice
	invoicesGroup.GET("/:id", controller.Set[dto.ByID](show))

	// Delete invoice
	invoicesGroup.DELETE("/:id", controller.Set[dto.ByID](deleteInvoice))
}
