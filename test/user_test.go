package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnums(t *testing.T) {
	newTest(t, httptest.NewRequest(http.MethodGet, "/user/test_di", nil)).
		checkStatusCode().
		checkBody("hello3 result")
}
