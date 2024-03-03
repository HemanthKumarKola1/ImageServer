// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: schedules.sql

package db

import (
	"context"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO schedules(
  schedule_time,
  user_id,
  message
) VALUES (
  $1, $2, $3
) RETURNING schedule_time, user_id, message
`

type CreateScheduleParams struct {
	ScheduleTime int32  `json:"schedule_time"`
	UserID       string `json:"user_id"`
	Message      string `json:"message"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule, arg.ScheduleTime, arg.UserID, arg.Message)
	var i Schedule
	err := row.Scan(&i.ScheduleTime, &i.UserID, &i.Message)
	return i, err
}

const deleteSchedule = `-- name: DeleteSchedule :exec
DELETE FROM schedules
WHERE schedule_time = $1
`

func (q *Queries) DeleteSchedule(ctx context.Context, scheduleTime int32) error {
	_, err := q.db.ExecContext(ctx, deleteSchedule, scheduleTime)
	return err
}

const getSchedule = `-- name: GetSchedule :one
SELECT schedule_time, user_id, message FROM schedules
WHERE schedule_time= $1 LIMIT 1
`

func (q *Queries) GetSchedule(ctx context.Context, scheduleTime int32) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, getSchedule, scheduleTime)
	var i Schedule
	err := row.Scan(&i.ScheduleTime, &i.UserID, &i.Message)
	return i, err
}