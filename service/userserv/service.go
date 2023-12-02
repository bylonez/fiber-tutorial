package userserv

import "fiber-tutorial/model"

func List() *[]*UserDTO {
	var userDtos []*UserDTO
	for _, user := range *model.List() {
		userDtos = append(userDtos, toDTO(user))
	}
	return &userDtos
}

func Create(u *UserCreateCmd) *UserDTO {
	user := model.Create(u.toUser())
	return toDTO(user)
}
