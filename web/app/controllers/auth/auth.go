package auth

import (
	"github.com/UchaBokeria/goyard/controller"
	"github.com/UchaBokeria/goyard/types"
)

func New(web *types.Goyard) {
	// HTMX endpoints for auth - no page handlers here
	route := web.Group("/auth")
	route.POST("/login", controller.Set[LoginDto](Login))
	route.POST("/register", controller.Set[RegisterDto](Register))
	route.POST("/logout", controller.Set[any](Logout))
}
