package main

import (
	"net/http"

	"github.com/Lev1reG/kairosai-backend/api"
	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/Lev1reG/kairosai-backend/db"
	"github.com/Lev1reG/kairosai-backend/internal/services"
	"github.com/Lev1reG/kairosai-backend/pkg/logger"
	"github.com/Lev1reG/kairosai-backend/pkg/utils"
	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()

	logger.InitLogger(cfg)
	logger.Log.Info("Starting KairosAI Backend", zap.String("environment", cfg.APP_ENV))

	database, err := db.ConnectDB(cfg)
	if err != nil {
		logger.Log.Fatal("Database connection failed", zap.Error(err))
		return
	}
	defer database.Close()

	if cfg.APP_ENV == "development" {
		err = db.RunMigrations(cfg)
		if err != nil {
			logger.Log.Fatal("Migration failed", zap.Error(err))
			return
		}
		logger.Log.Info("Migrations applied successfully")
	}

	authService := services.NewAuthService(database, cfg.JWT_SECRET)
	scheduleService := services.NewScheduleService(database, cfg.JWT_SECRET)

	handlers := &api.Handlers{
		AuthHandler:     api.NewAuthHandler(authService),
		ScheduleHandler: api.NewScheduleHandler(scheduleService),
	}

	utils.InitOAuth(cfg)

	r := api.SetupRoutes(handlers)

	port := cfg.PORT
	logger.Log.Info("Server running", zap.String("port", port))
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.Log.Fatal("Failed to start server: ", zap.Error(err))
	}
}
