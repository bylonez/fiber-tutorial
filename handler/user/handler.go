package user

import (
	"fiber-tutorial/common"
	"fiber-tutorial/model"
	"fiber-tutorial/service/userserv"
	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		userQuery := &model.UserQuery{}
		common.Parse(userQuery, c.QueryParser)
		return c.JSON(&common.Result{Data: userserv.List(userQuery)})
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
