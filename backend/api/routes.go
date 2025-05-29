package api

import (
	"net/http"

	"github.com/Lev1reG/kairosai-backend/api/middlewares"
	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Handlers struct {
	AuthHandler     *AuthHandler
	ScheduleHandler *ScheduleHandler
	// Add more handlers here (e.g., UserHandler, ScheduleHandler)
}

func SetupRoutes(handlers *Handlers) *chi.Mux {
	r := chi.NewRouter()
	cfg := config.LoadConfig()

	var allowedOrigins []string
	switch cfg.APP_ENV {
	case "development":
		allowedOrigins = []string{"http://localhost:5173"}
	default:
		allowedOrigins = []string{"https://kairos-ai-tau.vercel.app"}
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.Use(middlewares.LoggingMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Welcome to KairosAI Backend Service")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Post("/api/chat", HandleChatWithAI)

	// tambahkan ini dari branch teman
	r.Mount("/api/auth", authRoutes(handlers.AuthHandler))
	r.Mount("/api/schedules", scheduleRoutes(handlers.ScheduleHandler))

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

	r.Post("/resend-verification", authHandler.ResendVerificationEmail)
	r.Post("/verify-email", authHandler.VerifyEmail)
	r.Post("/forgot-password", authHandler.RequestResetPassword)
	r.Post("/reset-password", authHandler.ResetPassword)

	r.Route("/oauth", func(r chi.Router) {
		r.Get("/{provider}/login", authHandler.RedirectToOAuthProvider)
		r.Get("/{provider}/callback", authHandler.OAuthLogin)
	})
	return r
}

func scheduleRoutes(scheduleHandler *ScheduleHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		middlewares.JWTMiddleware(http.HandlerFunc(scheduleHandler.GetAllSchedules)).ServeHTTP(w, r)
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		middlewares.JWTMiddleware(http.HandlerFunc(scheduleHandler.GetScheduleDetail)).ServeHTTP(w, r)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		middlewares.JWTMiddleware(http.HandlerFunc(scheduleHandler.CreateSchedules)).ServeHTTP(w, r)
	})
	r.Delete("/{id}/cancel", func(w http.ResponseWriter, r *http.Request) {
		middlewares.JWTMiddleware(http.HandlerFunc(scheduleHandler.CancelSchedule)).ServeHTTP(w, r)
	})
	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		middlewares.JWTMiddleware(http.HandlerFunc(scheduleHandler.UpdateSchedule)).ServeHTTP(w, r)
	})

	return r
}
