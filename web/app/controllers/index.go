package controllers

import (
	Interface "main/web/app/types"
	"main/web/app/view/pages"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/UchaBokeria/goyard/types"
)

func Pages(web *types.Goyard) {
	// Main page routes - all SSR with inline functions
	web.GET("", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.Landing(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
	web.GET("/", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.Landing(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
	web.GET("/about", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.About(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
	web.GET("/checkout", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.Checkout(ctx.Get("webconfig").(Interface.WebConfig)))
	}))

	// Products pages - just render the page, HTMX will handle data loading
	web.GET("/products", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.Products(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
	web.GET("/products/:id", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.ProductDetail(ctx.Get("webconfig").(Interface.WebConfig)))
	}))

	// Auth pages
	web.GET("/login", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.Auth(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
	web.GET("/register", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.Auth(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
	web.GET("/forgot-password", controller.Set[any](func(ctx *controller.Context[any]) error {
		return ctx.Html(pages.ForgotPassword(ctx.Get("webconfig").(Interface.WebConfig)))
	}))
}
