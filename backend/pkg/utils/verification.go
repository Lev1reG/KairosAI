package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateSecureToken() (string, error) {
  bytes := make([]byte, 32)
  _, err := rand.Read(bytes)
  if err != nil {
    return "", err
  }
  return base64.URLEncoding.EncodeToString(bytes), nil
}

func HashToken(token string) string {
  hashed := sha256.Sum256([]byte(token))
  return base64.URLEncoding.EncodeToString(hashed[:])
}
