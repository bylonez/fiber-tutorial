package user

import (
	"fiber-tutorial/common"
	"fiber-tutorial/common/enum"
	"fiber-tutorial/model"
	"fiber-tutorial/service/servicei"
	"fiber-tutorial/service/userserv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"log"
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

	router.Get("/export", func(c *fiber.Ctx) error {
		userQuery := &model.UserQuery{}
		common.Parse(userQuery, c.QueryParser)
		userDtos := userserv.List(userQuery)
		var vos []userExportVO
		for _, userDTO := range *userDtos {
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
		formFile, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}

		file, err := formFile.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		xlsx, err := excelize.OpenReader(file)
		if err != nil {
			log.Fatal(err)
		}

		sheets := xlsx.GetSheetMap()
		for _, sheetName := range sheets {
			rows, err := xlsx.GetRows(sheetName)
			if err != nil {
				return err
			}

			for _, row := range rows {
				for _, cell := range row {
					fmt.Printf("%s\t", cell)
				}
				fmt.Println()
			}
		}

		return c.JSON(&common.Result{})
	})

}
