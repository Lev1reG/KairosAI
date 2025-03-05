package utils

import (
	"net/http"
	"time"

	"github.com/Lev1reG/kairosai-backend/config"
)

func SetAuthCookie(w http.ResponseWriter, token string) {
	cfg := config.LoadConfig()
	if cfg.APP_ENV == "development" {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
		})
	}
}

func ClearAuthCookie(w http.ResponseWriter) {
	cfg := config.LoadConfig()
	if cfg.APP_ENV == "development" {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    "",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(-time.Hour),
			MaxAge:   -1,
			Path:     "/",
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    "",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Expires:  time.Now().Add(-time.Hour),
			MaxAge:   -1,
			Path:     "/",
		})
	}
}
