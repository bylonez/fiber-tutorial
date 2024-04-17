package test

import (
	"fiber-tutorial/service"
)

type TestServiceImpl struct {
}

func (t TestServiceImpl) Hello2() string {
	return service.UserService.Hello3()
}

func init() {
	service.TestService = TestServiceImpl{}
}
