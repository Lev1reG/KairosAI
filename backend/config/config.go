package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV    string
	PORT       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		APP_ENV:    getEnv("APP_ENV", "development"),
		PORT:       getEnv("PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "db_user"),
		DBPassword: getEnv("DB_PASSWORD", "db_password"),
		DBName:     getEnv("DB_NAME", "db_name"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
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
