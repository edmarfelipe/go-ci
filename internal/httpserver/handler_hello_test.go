package httpserver_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edmarfelipe/go-ci/internal/httpserver"
)

func TestHelloHandler(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(httpserver.HelloHandler))

	t.Run("Should return a hello response with default language", func(t *testing.T) {
		resp, err := http.Get(srv.URL)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()

		var hello httpserver.HelloResponse
		json.NewDecoder(resp.Body).Decode(&hello)

		if hello.Message != "Hello" {
			t.Fatalf("expected message %s, got %s", "Hello", hello.Message)
		}
	})
}
