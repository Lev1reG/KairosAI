package db

import (
	"context"

	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var DB *pgxpool.Pool

func ConnectDB(cfg *config.Config) (*pgxpool.Pool, error) {
  connStr := cfg.GetDBConnectionString()

  pool, err := pgxpool.New(context.Background(), connStr)
  if err != nil {
    logger.Log.Fatal("Failed to connect to database: ", zap.Error(err))
    return nil, err
  }

  err = pool.Ping(context.Background())
  if err != nil {
    logger.Log.Fatal("Failed to ping database: ", zap.Error(err))
    return nil, err
  }

  logger.Log.Info("Connected to database")
  DB = pool
  return pool, nil
}
