package test

import (
	"encoding/json"
	"github.com/bylonez/fiber-tutorial/internal/handler"
	_ "github.com/bylonez/fiber-tutorial/internal/service/init"
	"github.com/bylonez/fiber-tutorial/pkg/dto"
	"github.com/gofiber/fiber/v3"
	"io"
	"net/http"
	"testing"
)

func setupApp() *fiber.App {
	app := fiber.New()
	handler.Route(app)
	// Define your routes or middleware here
	return app
}

type test struct {
	t    *testing.T
	resp *http.Response
}

func (t *test) checkErr(err error, msg string) {
	if err != nil {
		t.t.Fatalf(msg, err)
	}
}

func newTest(t *testing.T, req *http.Request) *test {
	app := setupApp()
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	return &test{
		t:    t,
		resp: resp,
	}
}

func (t *test) checkStatusCode() *test {
	if t.resp.StatusCode != http.StatusOK {
		t.t.Errorf("Expected status 200, got %v", t.resp.StatusCode)
	}
	return t
}

func (t *test) getBody() any {
	body, err := io.ReadAll(t.resp.Body)
	t.checkErr(err, "Expected no error reading response body, got %v")
	r := &dto.Result{}
	err = json.Unmarshal(body, r)
	t.checkErr(err, "Expected no error reading response body, got %v")
	return r.Data
}

func (t *test) checkBody(expected any) {
	body := t.getBody()
	if body != expected {
		t.t.Errorf("Expected body to be %v, got %v", expected, body)
	}
}
