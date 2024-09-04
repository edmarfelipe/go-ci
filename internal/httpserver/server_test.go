package httpserver_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edmarfelipe/go-ci/internal/env"
	"github.com/edmarfelipe/go-ci/internal/httpserver"
)

func TestHTTPServer(t *testing.T) {
	t.Run("Should start and stop the http server successfully", func(t *testing.T) {
		srv := httpserver.New(&env.Env{
			ServiceAddr: ":8080",
		})

		ts := httptest.NewServer(srv.Handler())
		defer ts.Close()

		errs := make(chan error, 1)
		go func() {
			err := srv.Start()
			errs <- err
		}()

		err := srv.Stop(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		err = <-errs
		if err != http.ErrServerClosed {
			t.Fatalf("expected %v, got %v", http.ErrServerClosed, err)
		}
	})
}
