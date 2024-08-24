package httpserver_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edmarfelipe/go-ci/internal/httpserver"
)

func TestHealthHandler(t *testing.T) {
	t.Run("Should return a health response", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(httpserver.HealthHandler))

		resp, err := http.Get(srv.URL)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		defer resp.Body.Close()

		var health httpserver.HealthResponse
		json.NewDecoder(resp.Body).Decode(&health)

		if health.Status != "ok" {
			t.Fatalf("expected status %s, got %s", "ok", health.Status)
		}
	})
}
