package controllers

import (
	"main/web/app/view/pages"

	"github.com/UchaBokeria/goyard/controller"
)

func Index(ctx *controller.Context[any]) error {
	return ctx.Html(pages.Landing())
}

func Products(ctx *controller.Context[any]) error {
	return ctx.Html(pages.Products())
}

func About(ctx *controller.Context[any]) error {
	return ctx.Html(pages.About())
}
