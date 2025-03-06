package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV            string
	APP_URL            string
	PORT               string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	JWT_SECRET         string
	GoogleClientID     string
	GoogleClientSecret string
	GithubClientID     string
	GithubClientSecret string
  SendGridAPIKey     string
	SMTPUsername       string
	SMTPPassword       string
	SMTPHost           string
	SMTPPort           string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		APP_ENV:            getEnv("APP_ENV", "development"),
		APP_URL:            getEnv("APP_URL", "http://localhost:8080"),
		PORT:               getEnv("PORT", "8080"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "db_user"),
		DBPassword:         getEnv("DB_PASSWORD", "db_password"),
		DBName:             getEnv("DB_NAME", "db_name"),
		DBSSLMode:          getEnv("DB_SSL_MODE", "disable"),
		JWT_SECRET:         getEnv("JWT_SECRET", "senpro2025"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", "google_client_id"),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", "google_client_secret"),
		GithubClientID:     getEnv("GITHUB_CLIENT_ID", "github_client_id"),
		GithubClientSecret: getEnv("GITHUB_CLIENT_SECRET", "github_client_secret"),
    SendGridAPIKey:     getEnv("SENDGRID_API_KEY", "sendgrid_api_key"),
		SMTPUsername:       getEnv("SMTP_USERNAME", "smtp_username"),
		SMTPPassword:       getEnv("SMTP_PASSWORD", "smtp_password"),
    SMTPHost:           getEnv("SMTP_HOST", "smtp_host"),
    SMTPPort:           getEnv("SMTP_PORT", "smtp_port"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (cfg *Config) GetDBConnectionString() string {
	return "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=" + cfg.DBSSLMode
}
