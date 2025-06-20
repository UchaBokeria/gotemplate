package admin

import "github.com/UchaBokeria/goyard/types"

func New(web *types.Goyard) {
	router := web.Group("/admin")
	router.Group("")
}
