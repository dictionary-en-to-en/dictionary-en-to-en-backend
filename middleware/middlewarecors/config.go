package middlewarecors

import "github.com/gofiber/fiber/v2/middleware/cors"

var (
	corsAllowMethods = "HEAD,GET,POST,PUT,DELETE,OPTIONS,PATCH"
	corsAllowOrigin  = "*"
	corsAllowHeaders = "*"

	Configs = cors.Config{
		Next:             nil,
		AllowOrigins:     corsAllowOrigin,
		AllowMethods:     corsAllowMethods,
		AllowHeaders:     corsAllowHeaders,
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
)
