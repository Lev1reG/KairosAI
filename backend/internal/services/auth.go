package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Lev1reG/kairosai-backend/db"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

	token, err := utils.GenerateSecureToken()
	if err != nil {
		logger.Log.Error("Error generating secure token", zap.Error(err))
		return nil, err
	}

	hashedToken := utils.HashToken(token)
	err = queries.CreateEmailVerificationToken(ctx, db.CreateEmailVerificationTokenParams{
		UserID: user.ID,
		Token:  hashedToken,
	})
	if err != nil {
		logger.Log.Error("Error inserting email verification token", zap.Error(err))
		return nil, errors.New("Failed to create email verification token")
	}

	if err := SendVerificationEmail(email, token); err != nil {
		logger.Log.Error("Error sending verification email", zap.Error(err))
		return nil, errors.New("Failed to send verification email")
	}

	logger.Log.Info("User registered", zap.String("user_id", user.ID.String()))
	return &user, nil
}

// Login a local user
func (a *AuthService) LoginUser(ctx context.Context, email, password string) (string, error) {
	queries := db.New(a.db)

	user, err := queries.GetVerifiedUserByEmail(ctx, email)
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

func (a *AuthService) GetUserByID(ctx context.Context, userID string) (*db.User, error) {
	queries := db.New(a.db)

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("Invalid user id format")
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	user, err := queries.GetUserByID(ctx, pgUUID)
	if err != nil {
		logger.Log.Error("Error getting user by ID", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

// OAuth Login
func (a *AuthService) OAuthLogin(ctx context.Context, provider, code string) (string, error) {
	queries := db.New(a.db)

	userInfo, err := utils.GetUserInfo(provider, code)
	if err != nil {
		return "", err
	}

	email, name, avatarURL, oauthID, err := utils.ExtractOAuthUserInfo(provider, userInfo)

	logger.Log.Debug("OAuth user info", zap.String("email", email), zap.String("name", name), zap.String("avatar_url", avatarURL), zap.String("oauth_id", oauthID))

	if email == "" || oauthID == "" {
		return "", errors.New("invalid user data received from OAuth provider")
	}

	existingUser, err := queries.GetUserByEmail(ctx, email)
	if err == nil {
		if existingUser.OauthProvider.String == "local" || existingUser.OauthProvider.String != provider {
			return "", errors.New("conflict: Email already registered with another provider")
		}

		return a.generateJWT(existingUser.ID.String())
	}

	username := generateUniqueUsername(ctx, queries, name)

	user, err := queries.CreateOAuthUser(ctx, db.CreateOAuthUserParams{
		Name:          name,
		Username:      username,
		Email:         email,
		AvatarUrl:     pgtype.Text{String: avatarURL, Valid: true},
		OauthProvider: pgtype.Text{String: provider, Valid: true},
		OauthID:       pgtype.Text{String: oauthID, Valid: true},
	})
	if err != nil {
		logger.Log.Error("Error inserting user", zap.Error(err))
		return "", err
	}

	return a.generateJWT(user.ID.String())
}

func (a *AuthService) VerifyEmail(ctx context.Context, token string) error {
	queries := db.New(a.db)

	hashedToken := utils.HashToken(token)
	userID, err := queries.GetUserByVerificationToken(ctx, hashedToken)
	if err != nil {
		logger.Log.Error("Error getting user by verification token", zap.Error(err))
		return errors.New("Invalid verification token")
	}

	err = queries.VerifyUserEmail(ctx, userID)
	if err != nil {
		logger.Log.Error("Error verifying user email", zap.Error(err))
		return errors.New("Failed to verify email")
	}

	_ = queries.DeleteEmailVerificationToken(ctx, userID)

	return nil
}

func generateUniqueUsername(ctx context.Context, queries *db.Queries, name string) string {
	baseUsername := sanitizeUsername(name)

	username := baseUsername
	for {
		_, err := queries.GetUserByUsername(ctx, username)
		if err != nil {
			break
		}
		username = fmt.Sprintf("%s_%d", baseUsername, rand.Intn(10000))
	}

	return username
}

func sanitizeUsername(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ReplaceAll(name, ".", "")
	name = strings.ReplaceAll(name, "@", "")
	return name
}
