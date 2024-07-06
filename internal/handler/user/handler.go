package user

import (
	"fiber-tutorial/internal/service/userservice"
	"fiber-tutorial/pkg/dto"
	"fiber-tutorial/pkg/ex"
	"fiber-tutorial/pkg/excel"
	"fiber-tutorial/pkg/valid"
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func Handler(router fiber.Router) {
	router.Get("/", func(c fiber.Ctx) error {
		userQuery := &userservice.UserQuery{}
		valid.Parse(userQuery, c.Bind().Query)
		return c.JSON(&dto.Result{Data: userservice.Service.List(userQuery)})
	})
	router.Post("/", func(c fiber.Ctx) error {
		userCreateCmd := &userservice.UserCreateCmd{}
		valid.Parse(userCreateCmd, c.Bind().Body)
		user := userservice.Service.Create(userCreateCmd)
		return c.JSON(&dto.Result{Data: user})
	})
	router.Put("/", func(c fiber.Ctx) error {
		userCreateCmd := &userservice.UserUpdateCmd{}
		valid.Parse(userCreateCmd, c.Bind().Body)
		user := userservice.Service.Update(userCreateCmd)
		return c.JSON(&dto.Result{Data: user})
	})
	router.Get("/test_di", func(c fiber.Ctx) error {
		return c.JSON(&dto.Result{Data: userservice.Service.Hello()})
	})

	router.Get("/enums", func(c fiber.Ctx) error {
		type EnumResult struct {
			Name string
			Desc string
		}
		var result []EnumResult
		for _, statusEnum := range userservice.StatusEnums {
			result = append(result, EnumResult{Name: statusEnum.Name(), Desc: statusEnum.Desc()})
		}
		return c.JSON(&dto.Result{Data: result})
	})

	router.Get("/export", func(c fiber.Ctx) error {
		userQuery := &userservice.UserQuery{}
		valid.Parse(userQuery, c.Bind().Query)
		userDtos := userservice.Service.List(userQuery)
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
			ex.ExportEmptyData.Panic()
		}
		excel.WriteResponse(vos, c)
		return nil
	})

	router.Post("/import", func(c fiber.Ctx) error {
		for _, u := range excel.ParseExcel[userImportCO](c, "file") {
			fmt.Println(*u)
		}
		return nil
	})

}
