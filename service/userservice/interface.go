package userservice

import (
	"fiber-tutorial/model"
)

type service interface {
	Hello() string
	Hello3() string
	List(query *model.UserQuery) []*UserDTO
	Create(cmd *UserCreateCmd) *UserDTO
	Update(cmd *UserUpdateCmd) *UserDTO
}

var Service service
