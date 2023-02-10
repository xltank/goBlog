package api

import (
	"goBlog/mid"

	"github.com/gofiber/fiber/v2"
)

func init() {
}

// Init api routes.
//
// Examples:
//	func InitRoutes(app) ...
func InitRoutes(app *fiber.App) {
	g := app.Group("/api", mid.AuthApi(nil))
	InitLoginRoutes(g)
	InitUserRoutes(g)
}
