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

	t.Run("Should return a hello response with language set to pt", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "?lang=pt")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()

		var hello httpserver.HelloResponse
		json.NewDecoder(resp.Body).Decode(&hello)

		if hello.Message != "Olá" {
			t.Fatalf("expected message %s, got %s", "Olá", hello.Message)
		}
	})

	t.Run("Should return error response with invalid language", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "?lang=invalid")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
		defer resp.Body.Close()

		var errResp httpserver.ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errResp)

		if errResp.Error != "invalid language" {
			t.Fatalf("expected error message %s, got %s", "invalid language", errResp.Error)
		}
	})
}
