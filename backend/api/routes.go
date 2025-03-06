package api

import (
	"net/http"

	"github.com/Lev1reG/kairosai-backend/api/middlewares"
	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	AuthHandler *AuthHandler
	// Add more handlers here (e.g., UserHandler, ScheduleHandler)
}

func SetupRoutes(handlers *Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Welcome to KairosAI Backend Service")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Mount("/api/auth", authRoutes(handlers.AuthHandler))

	return r
}

func authRoutes(authHandler *AuthHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
  r.Post("/logout", authHandler.Logout)

	r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		middlewares.JWTMiddleware(http.HandlerFunc(authHandler.GetCurrentUser)).ServeHTTP(w, r)
	})

  r.Get("/verify-email", authHandler.VerifyEmail)

  r.Route("/oauth", func(r chi.Router) {
    r.Get("/{provider}/login", authHandler.RedirectToOAuthProvider)
    r.Get("/{provider}/callback", authHandler.OAuthLogin)
  })
	return r
}
