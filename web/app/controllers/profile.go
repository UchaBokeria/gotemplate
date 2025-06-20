package controllers

import (
	"main/web/app/view/pages"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/a-h/templ"
)

func Profile(ctx *controller.Context[any]) error {
	return ctx.Html(pages.Profile())
}

func Orders(ctx *controller.Context[any]) error {
	return ctx.Html(templ.NopComponent)
}
