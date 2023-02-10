package api

import (
	"fmt"
	userservice "goBlog/service/userService"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitUserRoutes(r fiber.Router) {
	g := r.Group("/user")

	g.Get("/", func(c *fiber.Ctx) error {
		list := userservice.GetUserList()
		return c.JSON(fiber.Map{
			"rtn":  0,
			"err":  "",
			"data": list,
		})
	})

	g.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			panic("Invalid id:" + c.Params("id"))
		}
		fmt.Println()
		user := userservice.GetUserById(id)
		return c.JSON(fiber.Map{
			"rtn":  0,
			"err":  "",
			"data": user,
		})
	})

}
