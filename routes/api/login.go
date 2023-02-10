package api

import (
	"fmt"
	userservice "goBlog/service/userService"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func InitLoginRoutes(r fiber.Router) {
	r.Post("/singup", func(c *fiber.Ctx) error {
		list := userservice.GetUserList()
		return c.JSON(fiber.Map{
			"rtn":  0,
			"err":  "",
			"data": list,
		})
	})

	r.Get("/login", func(c *fiber.Ctx) error {
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
