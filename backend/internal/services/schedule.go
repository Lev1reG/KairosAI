package services

import (
	"context"
	"errors"

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

func NewScheduleService(db *pgxpool.Pool, jwtSecret string) *ScheduleService {
	return &ScheduleService{
		db:        db,
		jwtSecret: jwtSecret,
	}
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
