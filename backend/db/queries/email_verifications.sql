-- name: CreateEmailVerificationToken :exec
INSERT INTO email_verifications (user_id, token, expires_at)
VALUES ($1, $2, NOW() + INTERVAL '15 minutes');

-- name: GetUserByVerificationToken :one
SELECT user_id FROM email_verifications
WHERE token = $1 AND expires_at > NOW();

-- name: DeleteEmailVerificationToken :exec
DELETE FROM email_verifications
WHERE user_id = $1;
