package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
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
	ctx := context.Background()

	// Load your service account key JSON
	data, err := os.ReadFile("google-credentials.json")
	if err != nil {
		return "", fmt.Errorf("failed to read credentials file: %w", err)
	}

		// Generate token
	creds, err := google.CredentialsFromJSON(ctx, data, "https://www.googleapis.com/auth/dialogflow")
	if err != nil {
		return "", fmt.Errorf("failed to create credentials: %w", err)
	}

		token, err := creds.TokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}

		return token.AccessToken, nil
}