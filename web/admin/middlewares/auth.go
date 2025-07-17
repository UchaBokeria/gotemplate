package middlewares

import (
	"github.com/UchaBokeria/goyard/controller"
	"github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Use[any](func(ctx *controller.Context[any]) error {
			ctx.Set("x", 0)
			return next(ctx)
		})
	}
}
