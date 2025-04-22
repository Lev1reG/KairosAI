-- name: GetSchedulesByUser :many
SELECT * FROM schedules
WHERE user_id = $1
ORDER BY start_time;

-- name: GetScheduleByID :one
SELECT * FROM schedules
WHERE id = $1 AND user_id = $2;
