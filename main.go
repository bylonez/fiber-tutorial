package main

import (
	"fiber-tutorial/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

type (
	User struct {
		Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
		Age  int    `validate:"required,teener"`       // Required field, and client needs to implement our 'teener' tag format which we'll see later
	}
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

	app.Get("/", func(c *fiber.Ctx) error {
		data := User{
			Name: "Bob",
			Age:  20,
		}
		//common.SystemError.Panic("aaa", "B")
		//panic("ab")
		return c.JSON(common.Result{Data: data})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		user := &User{
			Name: c.Query("name"),
			Age:  c.QueryInt("age"),
		}
		common.Valid(user)
		return c.JSON(common.Result{Data: user})
	})

	log.Fatal(app.Listen(":3000"))
}
