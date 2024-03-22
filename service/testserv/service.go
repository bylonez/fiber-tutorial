package testserv

import (
	"fiber-tutorial/service/servicei"
)

type TestServiceImpl struct {
}

func (t TestServiceImpl) Hello2() string {
	return servicei.UserService.Hello3()
}

func init() {
	servicei.TestService = TestServiceImpl{}
}
