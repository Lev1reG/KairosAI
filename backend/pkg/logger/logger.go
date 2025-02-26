package logger

import (
	"github.com/Lev1reG/kairosai-backend/config"
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(cfg *config.Config) {
	var logger *zap.Logger
	var err error

	if cfg.APP_ENV == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	Log = logger
	defer Log.Sync()
}
