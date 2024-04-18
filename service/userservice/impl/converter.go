package impl

import "fiber-tutorial/service/userservice"

func createCmdToUser(c *userservice.UserCreateCmd) *User {
	return &User{
		Name:     c.Name,
		Birthday: c.Birthday,
		Gender:   c.Gender,
	}
}

func updateCmdToUser(c *userservice.UserUpdateCmd) *User {
	return &User{
		ID:       c.Id,
		Name:     c.Name,
		Birthday: c.Birthday,
		Gender:   c.Gender,
	}
}

func ToDTO(u *User) *userservice.UserDTO {
	return &userservice.UserDTO{
		Id:       u.ID,
		Name:     u.Name,
		Birthday: u.Birthday,
		Gender:   u.Gender,
	}
}

func ToDtos(users []*User) []*userservice.UserDTO {
	var userDtos []*userservice.UserDTO
	for _, user := range users {
		userDtos = append(userDtos, ToDTO(user))
	}
	return userDtos
}
