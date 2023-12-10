package userserv

import (
	"fiber-tutorial/common/field"
	"fiber-tutorial/model"
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
)

func (c *UserCreateCmd) toUser() *model.User {
	return &model.User{
		Name:     c.Name,
		Birthday: c.Birthday,
		Gender:   c.Gender,
	}
}

func (c *UserUpdateCmd) toUser() *model.User {
	return &model.User{
		ID:       c.Id,
		Name:     c.Name,
		Birthday: c.Birthday,
		Gender:   c.Gender,
	}
}

func toDTO(u *model.User) *UserDTO {
	return &UserDTO{
		Id:       u.ID,
		Name:     u.Name,
		Birthday: u.Birthday,
		Gender:   u.Gender,
	}
}

func toDtos(users *[]*model.User) *[]*UserDTO {
	var userDtos []*UserDTO
	for _, user := range *users {
		userDtos = append(userDtos, toDTO(user))
	}
	return &userDtos
}
