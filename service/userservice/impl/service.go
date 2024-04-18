package impl

import (
	"fiber-tutorial/model"
	"fiber-tutorial/service/testservice"
	"fiber-tutorial/service/userservice"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) Hello3() string {
	return "hello3 result"
}

func (u UserServiceImpl) Hello() string {
	return testservice.TestService.Hello2()
}

func (u UserServiceImpl) List(query *model.UserQuery) []*userservice.UserDTO {
	return userservice.ToDtos(model.ListUser(query))
}

func (u UserServiceImpl) Create(cmd *userservice.UserCreateCmd) *userservice.UserDTO {
	user := model.CreateUser(cmd.ToUser())
	return userservice.ToDTO(user)
}

func (u UserServiceImpl) Update(cmd *userservice.UserUpdateCmd) *userservice.UserDTO {
	user := model.UpdateUser(cmd.ToUser())
	return userservice.ToDTO(user)
}

func init() {
	userservice.UserService = UserServiceImpl{}
}
