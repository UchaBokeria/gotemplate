package app

import (
	"main/web/app/controllers"
	"main/web/app/controllers/products"
	"main/web/app/middlewares"

	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	// Apply translation middleware globally
	web.Use(middlewares.WebConfig())

	// Main page routes
	controllers.Pages(web)

	// Feature controllers for HTMX requests
	products.New(web)
}
