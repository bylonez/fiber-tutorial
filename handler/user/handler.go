package user

import (
	"fiber-tutorial/common"
	"fiber-tutorial/model"
	"fiber-tutorial/service/userserv"
	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&common.Result{Data: userserv.List()})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		user := &User{}
		err := c.BodyParser(user)
		if err != nil {
			return err
		}

		common.Valid(user)
		realUser := model.User{
			Name: user.Name,
		}
		userserv.Create(&realUser)
		return c.JSON(&common.Result{Data: &realUser})
	})
}
