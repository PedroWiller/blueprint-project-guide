package tests

import (
	"io"
	"net/http"
	"testing"

	"blueprint-project/internal/server"

	"github.com/gofiber/fiber/v2"
)

func TestHandler(t *testing.T) {
	app := fiber.New()

	s := &server.FiberServer{App: app}

	app.Get("/", s.HelloWorldHandler)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("error creating request. Err: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "<h1>Welcome to blueprint project guide!</h1>"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}
