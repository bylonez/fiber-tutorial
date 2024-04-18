package impl

import (
	"fiber-tutorial/service/testservice"
	"fiber-tutorial/service/userservice"
)

type service struct{}

func init() {
	userservice.Service = &service{}
}

func (s *service) Hello3() string {
	return "hello3 result"
}

func (s *service) Hello() string {
	return testservice.TestService.Hello2()
}

func (s *service) List(query *userservice.UserQuery) []*userservice.UserDTO {
	return ToDtos(ListUser(query))
}

func (s *service) Create(cmd *userservice.UserCreateCmd) *userservice.UserDTO {
	user := CreateUser(createCmdToUser(cmd))
	return ToDTO(user)
}

func (s *service) Update(cmd *userservice.UserUpdateCmd) *userservice.UserDTO {
	user := UpdateUser(updateCmdToUser(cmd))
	return ToDTO(user)
}
