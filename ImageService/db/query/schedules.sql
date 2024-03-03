-- name: CreateSchedule :one
INSERT INTO schedules(
  schedule_time,
  user_id,
  message
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetSchedule :one
SELECT * FROM schedules
WHERE schedule_time= $1 LIMIT 1;

-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE schedule_time = $1;