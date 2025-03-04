package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Lev1reG/kairosai-backend/internal/services"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
)

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

	utils.SuccessResponse(w, http.StatusCreated, "User registered successfully", user)
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
