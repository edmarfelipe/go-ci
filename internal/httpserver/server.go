package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/edmarfelipe/go-ci/internal/env"
)

type HTTPServer struct {
	srv *http.Server
}

func New(cfg *env.Env) *HTTPServer {
	router := http.NewServeMux()
	router.HandleFunc("GET /health", HealthHandler)
	router.HandleFunc("GET /hello", HelloHandler)

	return &HTTPServer{
		srv: &http.Server{
			Addr:    cfg.ServiceAddr,
			Handler: router,
		},
	}
}

func (s *HTTPServer) Start() error {
	slog.Info("starting http server", "addr", s.srv.Addr)
	return s.srv.ListenAndServe()
}
