package userservice

import (
	"fiber-tutorial/model"
)

type UserServiceI interface {
	Hello() string
	Hello3() string
	List(query *model.UserQuery) []*UserDTO
	Create(cmd *UserCreateCmd) *UserDTO
	Update(cmd *UserUpdateCmd) *UserDTO
}

var UserService UserServiceI
