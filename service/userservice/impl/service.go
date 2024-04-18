package impl

import (
	"fiber-tutorial/model"
	"fiber-tutorial/service/testservice"
	"fiber-tutorial/service/userservice"
)

type service struct {
}

func (s service) Hello3() string {
	return "hello3 result"
}

func (s service) Hello() string {
	return testservice.TestService.Hello2()
}

func (s service) List(query *model.UserQuery) []*userservice.UserDTO {
	return userservice.ToDtos(model.ListUser(query))
}

func (s service) Create(cmd *userservice.UserCreateCmd) *userservice.UserDTO {
	user := model.CreateUser(cmd.ToUser())
	return userservice.ToDTO(user)
}

func (s service) Update(cmd *userservice.UserUpdateCmd) *userservice.UserDTO {
	user := model.UpdateUser(cmd.ToUser())
	return userservice.ToDTO(user)
}

func init() {
	userservice.Service = service{}
}
