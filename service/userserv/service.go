package userserv

import "fiber-tutorial/model"

func List() *[]model.User {
	return model.List()
}

func Create(user *model.User) {
	model.Create(user)
}
