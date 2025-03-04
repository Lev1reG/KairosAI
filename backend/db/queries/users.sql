-- name: CreateUser :one
INSERT INTO users (name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, 'local', NULL, NULL, NOW(), NOW())
RETURNING id, name, username, email, created_at, updated_at;

-- name: CreateOAuthUser :one
INSERT INTO users (name, username, email, oauth_provider, oauth_id, avatar_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
ON CONFLICT (email) DO UPDATE
SET oauth_provider = EXCLUDED.oauth_provider,
    oauth_id = EXCLUDED.oauth_id,
    avatar_url = EXCLUDED.avatar_url,
    updated_at = NOW()
RETURNING id, name, username, email, avatar_url, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT id, name, username, email, password_hash, oauth_provider, oauth_id, avatar_url, created_at, updated_at
FROM users
WHERE username = $1;

-- name: GetUserByOAuthID :one
SELECT id, name, username, email, oauth_provider, oauth_id, avatar_url, created_at, updated_at
FROM users
WHERE oauth_provider = $1 AND oauth_id = $2;
