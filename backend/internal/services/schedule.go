package services

import (
	"context"
	"errors"
	"time"

	"github.com/Lev1reG/kairosai-backend/db"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ScheduleService struct {
	db        *pgxpool.Pool
	jwtSecret string
}

type CreateScheduleInput struct {
	Title       string     `json:"title" validate:"required,min=3,max=100"`
	Description *string    `json:"description,omitempty" validate:"omitempty,max=500"`
	StartTime   time.Time  `json:"start_time" validate:"required"`
	EndTime     *time.Time `json:"end_time,omitempty" validate:"omitempty,gtfield=StartTime"`
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

func (s *ScheduleService) CancelScheduleByID(ctx context.Context, userID string, scheduleID string) error {
	queries := db.New(s.db)

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("Invalid user ID format")
	}

	scheduleUUID, err := uuid.Parse(scheduleID)
	if err != nil {
		return errors.New("Invalid schedule ID format")
	}

	_, err = queries.GetNonCanceledSchedulesByID(ctx, db.GetNonCanceledSchedulesByIDParams{
		ID:     pgtype.UUID{Bytes: scheduleUUID, Valid: true},
		UserID: pgtype.UUID{Bytes: userUUID, Valid: true},
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("Schedule not found")
		}
		return err
	}

	err = queries.SoftDeleteScheduleByID(ctx, db.SoftDeleteScheduleByIDParams{
		ID:     pgtype.UUID{Bytes: scheduleUUID, Valid: true},
		UserID: pgtype.UUID{Bytes: userUUID, Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleService) CreateSchedule(ctx context.Context, input CreateScheduleInput) (*db.Schedule, error) {
	queries := db.New(s.db)

	parsedUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, errors.New("Invalid user id format")
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	endTime := input.EndTime
	if endTime == nil {
		defaultEnd := input.StartTime.Add(1 * time.Hour)
		endTime = &defaultEnd
	}

	if endTime.Before(input.StartTime) {
		return nil, errors.New("End time cannot be before start time")
	}

	params := db.CreateScheduleParams{
		UserID:      pgUUID,
		Title:       input.Title,
		Description: pgtype.Text{String: "", Valid: false},
		StartTime:   toTimestamptz(input.StartTime),
		EndTime:     toTimestamptz(*endTime),
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

	if endTime.Sub(input.StartTime) > 24*time.Hour {
		return nil, errors.New("Schedule duration cannot exceed 24 hours")
	}

	if input.StartTime.Before(time.Now()) {
		return nil, errors.New("Start time cannot be in the past")
	}

	schedule, err := queries.CreateSchedule(ctx, params)
	if err != nil {
		logger.Log.Error("Failed to create schedule", zap.Error(err))
		return nil, err
	}

	return &schedule, nil
}

func (s *ScheduleService) GetSchedulesByUser(ctx context.Context, userID string, limit, offset int32) (*[]db.Schedule, error) {
	queries := db.New(s.db)

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("Invalid user id format")
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	schedules, err := queries.GetSchedulesByUserWithPagination(ctx, db.GetSchedulesByUserWithPaginationParams{
		UserID: pgUUID,
		Limit:  limit,
		Offset: offset,
	})
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

func (s *ScheduleService) CountSchedulesByUser(ctx context.Context, userID string) (int64, error) {
	queries := db.New(s.db)

	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return 0, errors.New("Invalid user id format")
	}

	pgUUID := pgtype.UUID{Bytes: parsedUUID, Valid: true}

	count, err := queries.CountSchedulesByUser(ctx, pgUUID)
	if err != nil {
		logger.Log.Error("Failed to count schedules by user", zap.Error(err))
		return 0, err
	}

	return count, nil
}

func (s *ScheduleService) GetScheduleDetail(ctx context.Context, userID string, scheduleID string) (*db.Schedule, error) {
	queries := db.New(s.db)

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("Invalid user id format")
	}

	scheduleUUID, err := uuid.Parse(scheduleID)
	if err != nil {
		return nil, errors.New("Invalid schedule id format")
	}

	schedule, err := queries.GetScheduleByID(ctx, db.GetScheduleByIDParams{
		ID:     pgtype.UUID{Bytes: scheduleUUID, Valid: true},
		UserID: pgtype.UUID{Bytes: userUUID, Valid: true},
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		logger.Log.Error("Failed to get schedule detail", zap.Error(err))
		return nil, err
	}

	return &schedule, nil
}
