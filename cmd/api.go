package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/edmarfelipe/go-ci/internal/env"
	"github.com/edmarfelipe/go-ci/internal/httpserver"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to start", "err", err)
		os.Exit(1)
	}
}

func run() error {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	cfg, err := env.Load()
	// if err != nil {
	// 	return fmt.Errorf("error loading configs: %w", err)
	// }

	err = httpserver.New(cfg).Start()
	if err != nil {
		return fmt.Errorf("error starting the http server: %w", err)
	}

	return nil
}
