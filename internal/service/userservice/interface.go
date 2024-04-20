package userservice

type service interface {
	Hello() string
	Hello3() string
	List(query *UserQuery) []*UserDTO
	Create(cmd *UserCreateCmd) *UserDTO
	Update(cmd *UserUpdateCmd) *UserDTO
}

var Service service
