package userserv

import (
	"fiber-tutorial/model"
	"time"
)

type (
	UserDTO struct {
		Id       int
		Name     string
		Birthday time.Time
		Gender   string
	}

	UserCreateCmd struct {
		Name     string `validate:"required,min=3,max=20"`
		Birthday time.Time
		Gender   string
	}
)

func (c *UserCreateCmd) toUser() *model.User {
	return &model.User{
		Name:     c.Name,
		Birthday: c.Birthday,
		Gender:   c.Gender,
	}
}
