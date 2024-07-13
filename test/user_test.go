package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnums(t *testing.T) {
	test := newTest(t, httptest.NewRequest(http.MethodGet, "/user/test_di", nil))
	test.checkStatusCode()
	body := test.getBody()
	expected := "hello3 result"
	if body != expected {
		t.Errorf("Expected body to be %v, got %v", expected, body)
	}
}
