package mid

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthApi(config map[string]any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println(c.IP(), c.Hostname(), c.OriginalURL())
		c.Locals("user", fiber.Map{"name": "Michael", "role": "user"})
		return c.Next()
	}
}

func AuthAdmin(config map[string]any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println(c.IP(), c.Hostname(), c.OriginalURL())
		c.Locals("user", fiber.Map{"name": "Jason", "role": "admin"})
		return c.Next()
	}
}
