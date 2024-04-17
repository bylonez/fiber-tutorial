package user

import (
	"fiber-tutorial/common"
	"fiber-tutorial/common/enum"
	"fiber-tutorial/model"
	"fiber-tutorial/service"
	"fiber-tutorial/service/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		userQuery := &model.UserQuery{}
		common.Parse(userQuery, c.QueryParser)
		return c.JSON(&common.Result{Data: user.List(userQuery)})
	})
	router.Post("/", func(c *fiber.Ctx) error {
		userCreateCmd := &user.UserCreateCmd{}
		common.Parse(userCreateCmd, c.BodyParser)
		user := user.Create(userCreateCmd)
		return c.JSON(&common.Result{Data: user})
	})
	router.Put("/", func(c *fiber.Ctx) error {
		userCreateCmd := &user.UserUpdateCmd{}
		common.Parse(userCreateCmd, c.BodyParser)
		user := user.Update(userCreateCmd)
		return c.JSON(&common.Result{Data: user})
	})
	router.Get("/test_di", func(c *fiber.Ctx) error {
		return c.JSON(&common.Result{Data: service.UserService.Hello()})
	})

	router.Get("/enums", func(c *fiber.Ctx) error {
		type EnumResult struct {
			Name string
			Desc string
		}
		var result []EnumResult
		for _, statusEnum := range enum.StatusEnums {
			result = append(result, EnumResult{Name: statusEnum.Name(), Desc: statusEnum.Desc()})
		}
		return c.JSON(&common.Result{Data: result})
	})

	router.Get("/export", func(c *fiber.Ctx) error {
		userQuery := &model.UserQuery{}
		common.Parse(userQuery, c.QueryParser)
		userDtos := user.List(userQuery)
		var vos []userExportVO
		for _, userDTO := range userDtos {
			vos = append(vos, userExportVO{
				Id:       userDTO.Id,
				Name:     userDTO.Name,
				Birthday: userDTO.Birthday,
				Gender:   userDTO.Gender,
			})
		}
		if vos == nil {
			common.ExportEmptyData.Panic()
		}
		common.WriteResponse(vos, c)
		return nil
	})

	router.Post("/import", func(c *fiber.Ctx) error {
		for _, u := range common.ParseExcel[userImportCO](c, "file") {
			fmt.Println(*u)
		}
		return nil
	})

}
