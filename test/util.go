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

func newTest(t *testing.T, req *http.Request) *test {
	app := setupApp()
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	test := &test{
		t:    t,
		resp: resp,
	}
	return test
}

func (t *test) checkStatusCode() {
	if t.resp.StatusCode != http.StatusOK {
		t.t.Errorf("Expected status 200, got %v", t.resp.StatusCode)
	}
}

func (t *test) getBody() any {
	body, err := io.ReadAll(t.resp.Body)
	if err != nil {
		t.t.Fatalf("Expected no error reading response body, got %v", err)
	}
	r := &dto.Result{}
	err = json.Unmarshal(body, r)
	if err != nil {
		t.t.Fatalf("Expected no error reading response body, got %v", err)
	}
	return r.Data
}
