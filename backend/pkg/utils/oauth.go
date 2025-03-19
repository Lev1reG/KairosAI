package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

// Initialize OAuth configuration
func InitOAuth(cfg *config.Config) {
	googleOauthConfig = &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.FE_URL + "/auth/google/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// Get OAuth2 configuration
func GetOAuthConfig(provider string) (*oauth2.Config, error) {
	switch provider {
	case "google":
		return googleOauthConfig, nil
	default:
		return nil, errors.New("Invalid OAuth provider")
	}
}

// Exchange OAuth code for user info
func GetUserInfo(provider string, code string) (map[string]interface{}, error) {
	config, err := GetOAuthConfig(provider)
	if err != nil {
		return nil, err
	}

	logger.Log.Info("Exchaning OAuth code", zap.String("provider", provider), zap.String("code", code))

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		logger.Log.Error("Failed to exchange OAuth token", zap.String("provider", provider), zap.String("code", code), zap.Error(err))
		return nil, errors.New("Failed to exchange token")
	}

	var userInfoURL string
	switch provider {
	case "google":
		userInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
	default:
		return nil, errors.New("Invalid OAuth provider")
	}

	req, _ := http.NewRequest("GET", userInfoURL, nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		logger.Log.Error("Failed to get user info")
		return nil, errors.New("Failed to get user info")
	}
	defer res.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		logger.Log.Error("Failed to decode user info", zap.Error(err))
		return nil, errors.New("Failed to decode user info")
	}

	logger.Log.Debug("User info: ", zap.Any("user_info", userInfo))

	return userInfo, nil
}

func GenerateRandomState() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "M+tsqA4WTB4WaP/rJ/TmPA=="
	}
	return base64.URLEncoding.EncodeToString(b)
}

func ExtractOAuthUserInfo(provider string, userInfo map[string]interface{}) (string, string, string, string) {
	var email, name, avatarURL, oauthID string

	switch provider {
	case "google":
		email, _ = userInfo["email"].(string)
		name, _ = userInfo["name"].(string)
		avatarURL, _ = userInfo["picture"].(string)
		oauthID, _ = userInfo["id"].(string)
	default:
		return "", "", "", ""
	}

	return email, name, avatarURL, oauthID
}
