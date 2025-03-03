package services

import (
	"context"
	"errors"

	"github.com/Lev1reG/kairosai-backend/db"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
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
