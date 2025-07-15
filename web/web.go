package web

import (
	"main/web/app"

	"github.com/UchaBokeria/goyard/goyard"
	"github.com/labstack/echo/v4/middleware"
)

func New() {
	web := goyard.New()
	web.Static("/", "./public/")
	web.Pre(middleware.RemoveTrailingSlash())
	// web.Use(middleware.Logger())
	// web.Use(middleware.Recover())
	// web.Use(middleware.CORS())
	// web.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	// web.Use(middleware.Gzip())
	// web.Use(middleware.HTTPSRedirect())
	// web.Use(middleware.HTTPSWWWRedirect())
	// web.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{TokenLookup: "header:X-CSRF-Token"}))
	// web.Use(middleware.RequestID())
	// web.Use(middleware.Secure())
	// web.Use(middleware.BodyLimit("10M"))
	// web.Use(middleware.Gzip())
	// web.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	// 	XSSProtection:         "1; mode=block",
	// 	ContentTypeNosniff:    "nosniff",
	// 	XFrameOptions:         "DENY",
	// 	HSTSMaxAge:            31536000,
	// 	ContentSecurityPolicy: "default-src 'self'; img-src *; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'",
	// }))

	// admin.New(web)
	app.New(web)
	web.Logger.Fatal(web.Run(":3000"))
}
