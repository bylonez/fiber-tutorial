package user

import (
	"fiber-tutorial/common"
	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		data := User{
			Name: "Bob",
			Age:  20,
		}
		//common.SystemError.Panic("aaa", "B")
		//panic("ab")
		return c.JSON(common.Result{Data: data})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		user := &User{
			Name: c.Query("name"),
			Age:  c.QueryInt("age"),
		}
		common.Valid(user)
		return c.JSON(common.Result{Data: user})
	})
}
