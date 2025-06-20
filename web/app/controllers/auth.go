package controllers

import (
	"main/web/app/schemas"
	"main/web/app/view/pages"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/a-h/templ"
)

func Auth(ctx *controller.Context[any]) error {
	return ctx.Html(pages.Auth())
}

func Login(ctx *controller.Context[schemas.Login]) error {
	return ctx.Html(templ.NopComponent)
}

func Register(ctx *controller.Context[schemas.Register]) error {
	return ctx.Html(templ.NopComponent)
}
