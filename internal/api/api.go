package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/cassianobraz/encurtadorURLGo/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(store store.Store) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/url", func(r chi.Router) {
			r.Post("/shorten", handleShortenURL(store))
			r.Get("/{code}", handleGetShortenedURL(store))
		})
	})

	return r
}

func sendJSON(w http.ResponseWriter, resp apiResponse, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error ao fazer mashal de json", "error", err)
		sendJSON(w, apiResponse{Error: "something went wrong "}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error ao enviar a resposta", "error", err)
		return
	}
}
