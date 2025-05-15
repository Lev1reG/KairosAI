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

-- name: CheckScheduleConflict :one
SELECT EXISTS (
  SELECT 1 FROM schedules
  WHERE user_id = $1
    AND status = 'scheduled'
    AND (
      start_time, end_time
  ) OVERLAPS (
    $2::timestamptz, $3::timestamptz
  )
);

-- name: GetSchedulesByUserWithPagination :many
SELECT * FROM schedules
WHERE user_id = $1
ORDER BY start_time
LIMIT $2 OFFSET $3;

-- name: CountSchedulesByUser :one
SELECT COUNT(*) FROM schedules
WHERE user_id = $1;

-- name: SoftDeleteScheduleByID :exec
UPDATE schedules
SET status = 'canceled'
WHERE id = $1 AND user_id = $2 AND status != 'canceled';

-- name: GetNonCanceledSchedulesByID :one
SELECT * FROM schedules
WHERE id = $1 AND user_id = $2 AND status != 'canceled';

-- name: UpdateScheduleByID :exec
UPDATE schedules
SET
  title = COALESCE($1, title),
  description = COALESCE($2, description),
  start_time = COALESCE($3, start_time),
  end_time = COALESCE($4, end_time)
WHERE id = $5 AND user_id = $6;