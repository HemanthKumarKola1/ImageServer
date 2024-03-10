-- name: CreateSchedule :one
INSERT INTO schedules(
  schedule_time,
  user_id,
  message,
  image_data
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetSchedule :one
SELECT * FROM schedules
WHERE schedule_time= $1 LIMIT 1;

-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE schedule_time = $1;