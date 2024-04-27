package userservice

import (
	"fiber-tutorial/pkg/dto"
	"fiber-tutorial/pkg/field"
)

type (
	UserDTO struct {
		Id       uint
		Name     string
		Birthday field.Date
		Gender   string
	}

	UserCreateCmd struct {
		Name     string `validate:"required,min=3,max=20"`
		Birthday field.Date
		Gender   string
	}

	UserUpdateCmd struct {
		Id       uint
		Name     string `validate:"required,min=3,max=20"`
		Birthday field.Date
		Gender   string
	}

	UserQuery struct {
		dto.PageQuery
		Name string
	}
)
