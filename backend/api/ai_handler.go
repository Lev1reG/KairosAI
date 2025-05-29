package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/oauth2/google"
)

func HandleChatWithAI(w http.ResponseWriter, r *http.Request) {
	token, err := GetDialogflowAccessToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Forward the user's message to Dialogflow API
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("POST", "https://dialogflow.googleapis.com/v2/projects/kairos-vfqe/agent/sessions/test-session:detectIntent", r.Body)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to contact AI", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		logger.Log.Error("failed to copy response body", zap.Error(err))
	}
}

func GetDialogflowAccessToken() (string, error) {
	if err := ensureGoogleCredentials(); err != nil {
		return "", err
	}

	ctx := context.Background()
	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/dialogflow")
	if err != nil {
		return "", err
	}
	token, err := creds.TokenSource.Token()
	return token.AccessToken, err
}

func ensureGoogleCredentials() error {
	credsJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	if credsJSON == "" {
		return fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS_JSON is not set")
	}

	// Write to a secure, temp file
	path := filepath.Join(os.TempDir(), "google-credentials.json")
	if err := os.WriteFile(path, []byte(credsJSON), 0600); err != nil {
		return fmt.Errorf("failed to write credentials file: %w", err)
	}

	// Tell Google SDK where to look
	return os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", path)
}