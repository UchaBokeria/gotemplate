package app

import (
	"main/web/app/controllers"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	web.GET("", controller.Set[any](controllers.Index))
	web.GET("/", controller.Set[any](controllers.Index))
	web.GET("/products", controller.Set[any](controllers.Products))
	web.GET("/about", controller.Set[any](controllers.About))
	web.GET("/auth", controller.Set[any](controllers.Auth))
	web.GET("/login", controller.Set[any](controllers.Auth))
	web.GET("/register", controller.Set[any](controllers.Auth))
}
