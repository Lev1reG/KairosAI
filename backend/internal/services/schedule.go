package services

import (
	"context"
	"errors"
	"time"

	"github.com/Lev1reG/kairosai-backend/db"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ScheduleService struct {
	db        *pgxpool.Pool
	jwtSecret string
}

type CreateScheduleInput struct {
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description *string   `json:"description" validate:"omitempty,max=500"`
	StartTime   time.Time `json:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	UserID      string
}

func toTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func NewScheduleService(db *pgxpool.Pool, jwtSecret string) *ScheduleService {
	return &ScheduleService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

func (s *ScheduleService) CreateSchedule(ctx context.Context, input CreateScheduleInput) (*db.Schedule, error) {
	queries := db.New(s.db)

	parsedUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, errors.New("Invalid user id format")
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	params := db.CreateScheduleParams{
		UserID:      pgUUID,
		Title:       input.Title,
		Description: pgtype.Text{String: "", Valid: false},
		StartTime:   toTimestamptz(input.StartTime),
		EndTime:     toTimestamptz(input.EndTime),
	}

	if input.Description != nil {
		params.Description = pgtype.Text{String: *input.Description, Valid: true}
	}

	conflict, err := queries.CheckScheduleConflict(ctx, db.CheckScheduleConflictParams{
		UserID:  params.UserID,
		Column2: params.StartTime,
		Column3: params.EndTime,
	})
	if err != nil {
		logger.Log.Error("Failed to check schedule conflict", zap.Error(err))
		return nil, err
	}
	if conflict {
		return nil, errors.New("You already have a schedule in this time range")
	}

	schedule, err := queries.CreateSchedule(ctx, params)
	if err != nil {
		logger.Log.Error("Failed to create schedule", zap.Error(err))
		return nil, err
	}

	return &schedule, nil
}

func (s *ScheduleService) GetSchedulesByUser(ctx context.Context, userID string) (*[]db.Schedule, error) {
	queries := db.New(s.db)

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("Invalid user id format")
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	schedules, err := queries.GetSchedulesByUser(ctx, pgUUID)
	if err != nil {
		logger.Log.Error("Failed to get schedules by user", zap.Error(err))
		return nil, err
	}

	if schedules == nil {
		s := make([]db.Schedule, 0)
		return &s, nil
	}

	return &schedules, nil
}
