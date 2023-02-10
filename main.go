package main

import (
	"goBlog/config"
	"goBlog/db"
	"goBlog/routes/admin"
	"goBlog/routes/api"
	"goBlog/util/number"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.Init()
	db.InitMysql()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(csrf.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method} ${path} ${status} ${bytesSentF} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
		CustomTags: map[string]logger.LogFunc{
			"bytesSentF": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(number.FormatByte(len(c.Response().Body())))
			},
		},
	}))

	app.Static("/", "./public")

	app.Use(favicon.New(favicon.Config{
		File: "./public/imgs/favicon.ico",
	}))

	api.InitRoutes(app)
	admin.InitRoutes(app)

	app.Listen(":" + config.Conf.Port)
}
