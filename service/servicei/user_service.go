package servicei

type UserServiceI interface {
	Hello() string
	Hello3() string
}

var UserService UserServiceI
