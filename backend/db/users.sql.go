// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOAuthUser = `-- name: CreateOAuthUser :one
INSERT INTO users (name, username, email, oauth_provider, oauth_id, avatar_url, created_at, updated_at, email_verified)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), TRUE)
ON CONFLICT (email) DO UPDATE
SET oauth_provider = EXCLUDED.oauth_provider,
    oauth_id = EXCLUDED.oauth_id,
    avatar_url = EXCLUDED.avatar_url,
    updated_at = NOW()
RETURNING id, name, username, email, avatar_url, created_at, updated_at
`

type CreateOAuthUserParams struct {
	Name          string      `json:"name"`
	Username      string      `json:"username"`
	Email         string      `json:"email"`
	OauthProvider pgtype.Text `json:"oauth_provider"`
	OauthID       pgtype.Text `json:"oauth_id"`
	AvatarUrl     pgtype.Text `json:"avatar_url"`
}

type CreateOAuthUserRow struct {
	ID        pgtype.UUID        `json:"id"`
	Name      string             `json:"name"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	AvatarUrl pgtype.Text        `json:"avatar_url"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateOAuthUser(ctx context.Context, arg CreateOAuthUserParams) (CreateOAuthUserRow, error) {
	row := q.db.QueryRow(ctx, createOAuthUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.OauthProvider,
		arg.OauthID,
		arg.AvatarUrl,
	)
	var i CreateOAuthUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, 'local', NULL, NULL, NOW(), NOW())
RETURNING id, name, username, email, created_at, updated_at
`

type CreateUserParams struct {
	Name         string      `json:"name"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash pgtype.Text `json:"password_hash"`
}

type CreateUserRow struct {
	ID        pgtype.UUID        `json:"id"`
	Name      string             `json:"name"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, email_verified, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.OauthProvider,
		&i.OauthID,
		&i.AvatarUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, email_verified, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.OauthProvider,
		&i.OauthID,
		&i.AvatarUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByOAuthID = `-- name: GetUserByOAuthID :one
SELECT id, name, username, email, oauth_provider, oauth_id, avatar_url, created_at, updated_at
FROM users
WHERE oauth_provider = $1 AND oauth_id = $2
`

type GetUserByOAuthIDParams struct {
	OauthProvider pgtype.Text `json:"oauth_provider"`
	OauthID       pgtype.Text `json:"oauth_id"`
}

type GetUserByOAuthIDRow struct {
	ID            pgtype.UUID        `json:"id"`
	Name          string             `json:"name"`
	Username      string             `json:"username"`
	Email         string             `json:"email"`
	OauthProvider pgtype.Text        `json:"oauth_provider"`
	OauthID       pgtype.Text        `json:"oauth_id"`
	AvatarUrl     pgtype.Text        `json:"avatar_url"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetUserByOAuthID(ctx context.Context, arg GetUserByOAuthIDParams) (GetUserByOAuthIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByOAuthID, arg.OauthProvider, arg.OauthID)
	var i GetUserByOAuthIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.OauthProvider,
		&i.OauthID,
		&i.AvatarUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, email_verified, created_at, updated_at
FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.OauthProvider,
		&i.OauthID,
		&i.AvatarUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getVerifiedUserByEmail = `-- name: GetVerifiedUserByEmail :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, email_verified, created_at, updated_at
FROM users
WHERE email = $1 AND email_verified = TRUE
`

func (q *Queries) GetVerifiedUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getVerifiedUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.OauthProvider,
		&i.OauthID,
		&i.AvatarUrl,
		&i.EmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifyUserEmail = `-- name: VerifyUserEmail :exec
UPDATE users SET email_verified = TRUE
WHERE id = $1
`

func (q *Queries) VerifyUserEmail(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, verifyUserEmail, id)
	return err
}
