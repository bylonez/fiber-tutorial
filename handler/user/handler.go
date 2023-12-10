package user

import (
	"fiber-tutorial/common"
	"fiber-tutorial/service/userserv"
	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&common.Result{Data: userserv.List()})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		userCreateCmd := &userserv.UserCreateCmd{}
		common.Parse(userCreateCmd, c.BodyParser)
		user := userserv.Create(userCreateCmd)
		return c.JSON(&common.Result{Data: user})
	})

	router.Put("/", func(c *fiber.Ctx) error {
		userCreateCmd := &userserv.UserUpdateCmd{}
		common.Parse(userCreateCmd, c.BodyParser)
		user := userserv.Update(userCreateCmd)
		return c.JSON(&common.Result{Data: user})
	})
}
