package main

import (
	"fiber-tutorial/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

type User struct {
	Name string
	Age  uint8
}

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

	app.Get("/", func(c *fiber.Ctx) error {
		data := User{
			Name: "Bob",
			Age:  20,
		}
		//common.SystemError.Panic("aaa", "B")
		//panic("ab")
		return c.JSON(common.Result{Data: data})
	})

	log.Fatal(app.Listen(":3000"))
}
