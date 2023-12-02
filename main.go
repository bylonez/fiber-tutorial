package main

import (
	"fiber-tutorial/common"
	"fiber-tutorial/handler"
	"github.com/gofiber/fiber/v2"
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

	handler.SetupHandler(app)

	log.Fatal(app.Listen(":3000"))
}
