package userservice

import (
	"github.com/bylonez/fiber-tutorial/pkg/dto"
	"github.com/bylonez/fiber-tutorial/pkg/field"
	"github.com/bylonez/fiber-tutorial/pkg/valid"
	"github.com/go-playground/validator/v10"
)

type (
	UserDTO struct {
		Id       uint
		Name     string
		Birthday field.Date
		Gender   string
	}

	UserCreateCmd struct {
		Name     string `valid:"required,min=3,max=20"`
		Birthday field.Date
		Gender   string `valid:"required,gender"`
	}

	UserUpdateCmd struct {
		Id       uint
		Name     string `valid:"required,min=3,max=20"`
		Birthday field.Date
		Gender   string
	}

	UserQuery struct {
		dto.PageQuery
		Name string
	}
)

func init() {
	valid.Register("gender", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "male" || fl.Field().String() == "female"
	})
}
