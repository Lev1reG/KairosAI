package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Missing or invalid token")
			return
		}

		userID, err := validateJWT(tokenString)
		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Extract token from cookie
func extractToken(r *http.Request) string {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return ""
	}

	return cookie.Value
}

// Validate JWT token & return user ID
func validateJWT(tokenString string) (string, error) {
	cfg := config.LoadConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(cfg.JWT_SECRET), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
      logger.Log.Warn("Attemp to use expired token", zap.Error(err))
      return "", errors.New("Invalid token")
		}
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("Invalid token")
	}

	return userID, nil
}
