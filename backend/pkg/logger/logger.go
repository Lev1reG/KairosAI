package logger

import (
	"os"
	"time"

	"github.com/Lev1reG/kairosai-backend/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(cfg *config.Config) {
	var logLevel zapcore.Level
	if cfg.APP_ENV == "production" {
		logLevel = zapcore.InfoLevel
	} else {
		logLevel = zapcore.DebugLevel
	}

	logFile := "logs/app.log"
	_ = os.Mkdir("logs", 0755)

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		logLevel,
	)

  consoleCore := zapcore.NewCore(
    zapcore.NewConsoleEncoder(encoderConfig),  
    zapcore.AddSync(os.Stdout),
    logLevel,
  )

  core := zapcore.NewTee(fileCore, consoleCore)

  Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
  defer Log.Sync()
}
