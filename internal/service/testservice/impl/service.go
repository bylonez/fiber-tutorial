package impl

import (
	"fiber-tutorial/internal/service/testservice"
	"fiber-tutorial/internal/service/userservice"
)

type service struct{}

func init() {
	testservice.Service = &service{}
}

func (t service) Hello2() string {
	return userservice.Service.Hello3()
}
