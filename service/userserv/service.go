package userserv

import "fiber-tutorial/model"

func List() *[]model.User {
	return model.List()
}

func Create(u *UserCreateCmd) *model.User {
	user := model.Create(u.toUser())
	return user
}
