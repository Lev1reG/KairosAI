package db

import (
	"path/filepath"

	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func RunMigrations(cfg *config.Config) error {
	databaseUrl := cfg.GetDBConnectionString()

	migrationsPath, _ := filepath.Abs("db/migrations")

	m, err := migrate.New(
		"file://"+migrationsPath,
		databaseUrl,
	)
	if err != nil {
		logger.Log.Error("Failed to create migration instance", zap.Error(err))
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Log.Error("Failed to apply migrations", zap.Error(err))
		return err
	}

	return nil
}
