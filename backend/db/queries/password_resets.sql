-- name: CreatePasswordResetToken :exec
INSERT INTO password_resets (user_id, token, expires_at)
VALUES ($1, $2, NOW() + INTERVAL '15 minutes');

-- name: GetUserByResetToken :one
SELECT user_id FROM password_resets
WHERE token = $1 AND expires_at > NOW();

-- name: DeletePasswordResetToken :exec
DELETE FROM password_resets
WHERE user_id = $1;
