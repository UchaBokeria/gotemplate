package app

import (
	"main/web/app/controllers"

	"github.com/UchaBokeria/goyard/controller"
	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	web.GET("", controller.Set[any](controllers.Index))
}
