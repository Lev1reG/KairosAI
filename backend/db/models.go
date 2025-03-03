// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ChatLog struct {
	ID        pgtype.UUID        `json:"id"`
	UserID    pgtype.UUID        `json:"user_id"`
	Message   string             `json:"message"`
	Response  pgtype.Text        `json:"response"`
	Timestamp pgtype.Timestamptz `json:"timestamp"`
}

type Notification struct {
	ID         pgtype.UUID        `json:"id"`
	UserID     pgtype.UUID        `json:"user_id"`
	ScheduleID pgtype.UUID        `json:"schedule_id"`
	SentAt     pgtype.Timestamptz `json:"sent_at"`
}

type Schedule struct {
	ID          pgtype.UUID        `json:"id"`
	UserID      pgtype.UUID        `json:"user_id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	StartTime   pgtype.Timestamptz `json:"start_time"`
	EndTime     pgtype.Timestamptz `json:"end_time"`
	Status      string             `json:"status"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type ScheduleParticipant struct {
	ID             pgtype.UUID        `json:"id"`
	UserID         pgtype.UUID        `json:"user_id"`
	ScheduleID     pgtype.UUID        `json:"schedule_id"`
	Email          string             `json:"email"`
	ResponseStatus string             `json:"response_status"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	ID            pgtype.UUID        `json:"id"`
	Name          string             `json:"name"`
	Username      string             `json:"username"`
	Email         string             `json:"email"`
	PasswordHash  pgtype.Text        `json:"password_hash"`
	OauthProvider pgtype.Text        `json:"oauth_provider"`
	OauthID       pgtype.Text        `json:"oauth_id"`
	AvatarUrl     pgtype.Text        `json:"avatar_url"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}
