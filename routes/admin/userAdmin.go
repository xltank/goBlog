package admin

import "github.com/gofiber/fiber/v2"

func InitUserRoutes(r fiber.Router) {
	g := r.Group("/user")

	g.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"rtn": 0,
			"err": "",
			"data": fiber.Map{
				"list": []fiber.Map{
					{
						"name": "aa",
						"role": "admin",
					},
					{
						"name": "bb",
						"role": "admin",
					},
				},
				"test": "aaaaaaaaa",
			},
		})
	})

}
