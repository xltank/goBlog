package admin

import (
	"goBlog/mid"

	"github.com/gofiber/fiber/v2"
)

func init() {
}

// Init admin routes.
//
// Examples:
//	func InitRoutes(app) ...
func InitRoutes(r fiber.Router) {
	g := r.Group("/admin", mid.AuthAdmin(nil))

	InitUserRoutes(g)
}
