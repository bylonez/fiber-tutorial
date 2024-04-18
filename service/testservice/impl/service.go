package impl

import (
	"fiber-tutorial/service/testservice"
	"fiber-tutorial/service/userservice"
)

type TestServiceImpl struct {
}

func (t TestServiceImpl) Hello2() string {
	return userservice.Service.Hello3()
}

func init() {
	testservice.TestService = TestServiceImpl{}
}
