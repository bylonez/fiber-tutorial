package user

import (
	"fiber-tutorial/common"
	"fiber-tutorial/common/enum"
	"fiber-tutorial/model"
	"fiber-tutorial/service/servicei"
	"fiber-tutorial/service/userserv"
	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		userQuery := &model.UserQuery{}
		common.Parse(userQuery, c.QueryParser)
		return c.JSON(&common.Result{Data: userserv.List(userQuery)})
	})
	router.Get("/test_di", func(c *fiber.Ctx) error {
		return c.JSON(&common.Result{Data: servicei.UserService.Hello()})
	})

	router.Get("/enums", func(c *fiber.Ctx) error {
		type EnumResult struct {
			Name string
			Desc string
		}
		result := []EnumResult{}
		for _, statusEnum := range enum.StatusEnums {
			result = append(result, EnumResult{Name: statusEnum.Name(), Desc: statusEnum.Desc()})
		}
		return c.JSON(&common.Result{Data: result})
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
