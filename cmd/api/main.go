package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/cassianobraz/encurtadorURLGo/internal/api"
	"github.com/cassianobraz/encurtadorURLGo/internal/store"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(rdb)

	handler := api.NewHandler(store)

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
