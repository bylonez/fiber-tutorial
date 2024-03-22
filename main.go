package main

import (
	"fiber-tutorial/common"
	"fiber-tutorial/handler"
	"fiber-tutorial/model"
	_ "fiber-tutorial/service/servicedi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: common.ErrorHandler,
	})

	// middleware
	// recover from panic and log
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			// todo log stack trace
		},
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05Z07:00",
		TimeZone:   "Asia/Shanghai",
	}))

	handler.SetupHandler(app)
	model.InitDB()

	log.Fatal(app.Listen(":3000"))
}
