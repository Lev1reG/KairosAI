-- name: GetSchedulesByUser :many
SELECT * FROM schedules
WHERE user_id = $1
ORDER BY start_time;

-- name: GetScheduleByID :one
SELECT * FROM schedules
WHERE id = $1 AND user_id = $2;

-- name: CreateSchedule :one
INSERT INTO schedules (
  user_id, title, description, start_time, end_time, status, created_at, updated_at
)
VALUES (
  $1, $2, $3, $4, $5, DEFAULT, NOW(), NOW() 
)
RETURNING *;
