package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/cassianobraz/encurtadorURLGo/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	handler := api.NewHandler()

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		Handler:      handler,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
