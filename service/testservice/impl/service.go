package impl

import (
	"fiber-tutorial/service/testservice"
	"fiber-tutorial/service/userservice"
)

type service struct{}

func init() {
	testservice.Service = &service{}
}

func (t service) Hello2() string {
	return userservice.Service.Hello3()
}
