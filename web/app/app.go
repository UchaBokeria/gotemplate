package app

import (
	"main/web/app/controllers"
	"main/web/app/controllers/products"
	"main/web/app/middlewares"

	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	web.Use(middlewares.WebConfig())
	controllers.Pages(web)
	products.New(web)
}
