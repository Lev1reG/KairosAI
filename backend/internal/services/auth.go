package services

import (
	"context"
	"errors"
	"time"

	"github.com/Lev1reG/kairosai-backend/db"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AuthService struct {
	db        *pgxpool.Pool
	jwtSecret string
}

func NewAuthService(db *pgxpool.Pool, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

// Register a new local user
func (a *AuthService) RegisterUser(ctx context.Context, name, username, email, password string) (*db.CreateUserRow, error) {
	queries := db.New(a.db)

	if existingUser, err := queries.GetUserByEmail(ctx, email); err == nil && existingUser.ID.String() != "" {
		return nil, errors.New("duplicate key value")
	}

	if existingUser, err := queries.GetUserByUsername(ctx, username); err == nil && existingUser.ID.String() != "" {
		return nil, errors.New("duplicate key value")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		logger.Log.Error("Failed to hash password", zap.Error(err))
		return nil, err
	}

	user, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:         name,
		Username:     username,
		Email:        email,
		PasswordHash: pgtype.Text{String: hashedPassword, Valid: true},
	})
	if err != nil {
		logger.Log.Error("Error inserting user", zap.Error(err))
		return nil, err
	}

	logger.Log.Info("User registered", zap.String("user_id", user.ID.String()))
	return &user, nil
}

// Login a local user
func (a *AuthService) LoginUser(ctx context.Context, email, password string) (string, error) {
	queries := db.New(a.db)

	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		logger.Log.Error("Database error while retrieving user", zap.Error(err))
		return "", errors.New("Invalid email or password")
	}

	if !utils.ComparePassword(user.PasswordHash.String, password) {
		logger.Log.Warn("Failed login attempt", zap.String("email", email))
		return "", errors.New("Invalid email or password")
	}

	tokenString, err := a.generateJWT(user.ID.String())
	if err != nil {
		logger.Log.Error("Failed to generate JWT", zap.Error(err))
		return "", errors.New("Internal server error")
	}

	logger.Log.Info("User logged in successfully", zap.String("user_id", user.ID.String()))

	return tokenString, nil
}

// Generate JWT Token
func (a *AuthService) generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
