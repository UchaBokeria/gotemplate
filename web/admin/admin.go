package admin

import (
	"main/web/admin/controllers"
	"main/web/admin/controllers/invoices"
	"main/web/admin/controllers/orders"
	"main/web/admin/controllers/products"
	"main/web/admin/controllers/settings"

	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	router := web.Group("/admin")

	controllers.Pages(router)
	invoices.New(router)
	orders.New(router)
	products.New(router)
	settings.New(router)
}
