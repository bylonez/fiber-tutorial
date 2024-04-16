package userserv

import (
	"fiber-tutorial/model"
	"fiber-tutorial/service/servicei"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) Hello3() string {
	return "hello3 result"
}

func (u UserServiceImpl) Hello() string {
	return servicei.TestService.Hello2()
}

func List(query *model.UserQuery) []*UserDTO {
	return toDtos(model.ListUser(query))
}

func Create(u *UserCreateCmd) *UserDTO {
	user := model.CreateUser(u.toUser())
	return toDTO(user)
}

func Update(u *UserUpdateCmd) *UserDTO {
	user := model.UpdateUser(u.toUser())
	return toDTO(user)
}

func init() {
	servicei.UserService = UserServiceImpl{}
}
