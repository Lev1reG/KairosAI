package utils

import (
	"net/http"
	"time"
)

func SetAuthCookie(w http.ResponseWriter, token string) {
  http.SetCookie(w, &http.Cookie{
    Name: "auth_token",
    Value: token,
    HttpOnly: true,
    Secure: true,
    SameSite: http.SameSiteStrictMode,
    Expires: time.Now().Add(24 * time.Hour),
  })
}
