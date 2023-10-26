package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

type User struct {
	Name string
	Age  uint8
}

type Result struct {
	Code int
	Data any
	Msg  string
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(&Result{
				Code: fiber.StatusInternalServerError,
				Msg:  err.Error()},
			)
		},
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
		data := &User{
			Name: "Bob",
			Age:  20,
		}
		//panic(11)
		return c.JSON(&Result{Data: data})
	})

	log.Fatal(app.Listen(":3000"))
}
