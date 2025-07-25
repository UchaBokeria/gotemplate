package controllers

import (
	pages "main/web/admin/view/pages"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func Pages(router *echo.Group) {
	router.GET("/", controller.Set[any](func(ctx *controller.Context) error {
		return ctx.Html(pages.Products())
	}))

	router.GET("/products", controller.Set[any](func(ctx *controller.Context) error {
		return ctx.Html(pages.Products())
	}))
	router.GET("/categories", controller.Set[any](func(ctx *controller.Context) error {
		return ctx.Html(pages.Categories())
	}))

	router.GET("/orders", controller.Set[any](func(ctx *controller.Context) error {
		return ctx.Html(pages.Orders())
	}))

	router.GET("/invoices", controller.Set[any](func(ctx *controller.Context) error {
		return ctx.Html(pages.Invoices())
	}))

	router.GET("/settings", controller.Set[any](func(ctx *controller.Context) error {
		return ctx.Html(pages.Settings())
	}))
}
