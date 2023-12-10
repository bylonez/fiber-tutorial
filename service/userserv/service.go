package userserv

import "fiber-tutorial/model"

func List() *[]*UserDTO {
	return toDtos(model.ListUser())
}

func Create(u *UserCreateCmd) *UserDTO {
	user := model.CreateUser(u.toUser())
	return toDTO(user)
}

func Update(u *UserUpdateCmd) *UserDTO {
	user := model.UpdateUser(u.toUser())
	return toDTO(user)
}
