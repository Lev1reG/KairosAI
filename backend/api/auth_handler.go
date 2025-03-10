package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Lev1reG/kairosai-backend/api/middlewares"
	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/Lev1reG/kairosai-backend/internal/services"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type UserResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	OAuthProvider string    `json:"oauth_provider"`
	OAuthID       string    `json:"oauth_id,omitempty"`
	AvatarURL     string    `json:"avatar_url,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Name == "" || req.Username == "" || req.Email == "" || req.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "All fields are required")
		return
	}

	if !utils.IsValidEmail(req.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	if err := utils.IsValidPassword(req.Password); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authService.RegisterUser(r.Context(), req.Name, req.Username, req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			utils.ErrorResponse(w, http.StatusConflict, "User with this email or username already exists")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, "User registered successfully. Please check your email to verify your account.", user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	if !utils.IsValidEmail(req.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	token, err := h.authService.LoginUser(r.Context(), req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SetAuthCookie(w, token)

	utils.SuccessResponse(w, http.StatusOK, "Login successful", map[string]string{"token": token})
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request context
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok || userID == "" {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.authService.GetUserByID(r.Context(), userID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	userResponse := UserResponse{
		ID:            user.ID.String(),
		Name:          user.Name,
		Username:      user.Username,
		Email:         user.Email,
		OAuthProvider: user.OauthProvider.String,
		OAuthID:       user.OauthID.String,
		AvatarURL:     user.AvatarUrl.String,
		CreatedAt:     user.CreatedAt.Time,
		UpdatedAt:     user.UpdatedAt.Time,
	}

	utils.SuccessResponse(w, http.StatusOK, "User retrieved successfully", userResponse)
}

func (h *AuthHandler) RedirectToOAuthProvider(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	oauthConfig, err := utils.GetOAuthConfig(provider)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid OAuth provider")
		return
	}

	state := utils.GenerateRandomState()

	cfg := config.LoadConfig()

	if cfg.APP_ENV == "development" {
		http.SetCookie(w, &http.Cookie{
			Name:     "oauth_state",
			Value:    state,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
			Expires:  time.Now().Add(10 * time.Minute),
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "oauth_state",
			Value:    state,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			Expires:  time.Now().Add(10 * time.Minute),
		})
	}

	authUrl := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

	http.Redirect(w, r, authUrl, http.StatusFound)
}

func (h *AuthHandler) OAuthLogin(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	code := r.URL.Query().Get("code")

	if provider == "" || code == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	queryState := r.URL.Query().Get("state")
	if queryState == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Missing OAuth state")
		return
	}

	cookie, err := r.Cookie("oauth_state")
	if err != nil || cookie.Value != queryState {
		utils.ErrorResponse(w, http.StatusForbidden, "Invalid OAuth state")
		return
	}

	token, err := h.authService.OAuthLogin(r.Context(), provider, code)
	if err != nil {
		if strings.Contains(err.Error(), "conflict: Email") {
			logger.Log.Warn("OAuth email conflict", zap.String("error", err.Error()))
			utils.ErrorResponse(w, http.StatusConflict, "Email already registered with another provider")
			return
		}
		logger.Log.Error("Failed to login with OAuth", zap.Error(err))
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to login with OAuth")
		return
	}

	utils.SetAuthCookie(w, token)
	utils.SuccessResponse(w, http.StatusOK, "Login successful", map[string]string{"token": token})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	utils.ClearAuthCookie(w)

	utils.SuccessResponse(w, http.StatusOK, "Logout successful", nil)
}

func (h *AuthHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid token")
		return
	}

	err := h.authService.VerifyEmail(r.Context(), token)
	if err != nil {
		if err.Error() == "Invalid verification token" {
			utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Email verified successfully! You can now log in", nil)
}

func (h *AuthHandler) ResendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if !utils.IsValidEmail(req.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	err := h.authService.ResendVerificationEmail(r.Context(), req.Email)
	if err != nil {
		switch err {
		case services.ErrTooManyRequest:
			utils.ErrorResponse(w, http.StatusTooManyRequests, err.Error())
		case services.ErrUserNotFound:
			utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		case services.ErrAlreadyVerified:
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		case services.ErrInternalServer, services.ErrEmailFailed:
			utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		default:
			utils.ErrorResponse(w, http.StatusInternalServerError, "Unxpected error")
		}
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Verification email sent successfully", nil)
}

func (h *AuthHandler) RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if !utils.IsValidEmail(req.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid email format")
		return
	}

	err := h.authService.RequestResetPassword(r.Context(), req.Email)
	if err != nil {
		switch err {
		case services.ErrTooManyRequest:
			utils.ErrorResponse(w, http.StatusTooManyRequests, err.Error())
		case services.ErrUserNotFound:
			utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		case services.ErrNotVerified:
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		case services.ErrInternalServer, services.ErrEmailFailed:
			utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		case services.ErrOauthUser:
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		default:
			utils.ErrorResponse(w, http.StatusInternalServerError, "Unxpected error")
		}
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Reset password email sent successfully", nil)
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		NewPassword string `json:"new_password"`
	}

	token := r.URL.Query().Get("token")
	if token == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid token")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.IsValidPassword(req.NewPassword); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.authService.ResetPassword(r.Context(), token, req.NewPassword)
	if err != nil {
		if err.Error() == "Invalid token" {
			utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	utils.SuccessResponse(w, http.StatusOK, "Password reset successfully", nil)
}
