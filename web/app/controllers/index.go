package controllers

import (
	"main/internal/types/dto"
	"main/web/app/view/pages"

	Interface "main/web/app/types"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/UchaBokeria/goyard/types"
)

func Pages(router *types.Goyard) {
	router.GET("/translate/:lang", controller.Set[Interface.WebDto](func(ctx *controller.Context, web *Interface.WebDto) error {
		ctx.WriteCookie(controller.Cookie{Key: "lang", Value: web.Lang, Path: controller.Ptr("/"), MaxAge: controller.Ptr(3600)})
		return ctx.Redirect(302, web.Redirect)
	}))

	// Main page routes - all SSR with inline functions
	router.GET("", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.Landing(web.Config))
	}))
	router.GET("/", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.Landing(web.Config))
	}))
	router.GET("/about", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.About(web.Config))
	}))
	router.GET("/checkout", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.Checkout(web.Config))
	}))

	// Products pages - just render the page, HTMX will handle data loading
	router.GET("/products", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.Products(web.Config))
	}))
	router.GET("/products/:id", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.ProductDetail(web.Config))
	}))

	// Auth pages
	router.GET("/login", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.Auth(web.Config))
	}))
	router.GET("/register", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.Auth(web.Config))
	}))
	router.GET("/forgot-password", controller.Set[dto.WebConfig](func(ctx *controller.Context, web *dto.WebConfig) error {
		return ctx.Html(pages.ForgotPassword(web.Config))
	}))
}
