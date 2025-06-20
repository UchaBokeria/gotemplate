package controllers

import (
	"main/web/app/view/pages"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/a-h/templ"
)

func Index(ctx *controller.Context[any]) error {
	return ctx.Html(pages.Landing())
}

func Products(ctx *controller.Context[any]) error {
	return ctx.Html(templ.NopComponent)
}

