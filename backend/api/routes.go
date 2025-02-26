package api

import (
	"net/http"

	"github.com/Lev1reG/kairosai-backend/api/middlewares"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

  r.Use(middlewares.LoggingMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    if _, err := w.Write([]byte("Welcome to KairosAI Backend Service")); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
	})

	return r
}
