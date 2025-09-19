package util

import (
	config "task-management/internal/adapter/config"

	"github.com/gofiber/fiber/v2"

	"fmt"
	"os"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	LoggerInstance *zap.Logger
	logWriter      *lumberjack.Logger
	currentDate    string
	logMutex       sync.RWMutex
)

func updateLoggerIfNeeded() {
	today := time.Now().Format("2006-01-02")

	logMutex.Lock()
	defer logMutex.Unlock()

	if currentDate != today {
		currentDate = today
		filename := fmt.Sprintf("logs/%s.log", today)

		if logWriter != nil {
			logWriter.Close()
		}

		logWriter = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    10,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   true,
			LocalTime:  true,
		}

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "time"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder := zapcore.NewConsoleEncoder(encoderCfg)

		var level zapcore.Level
		switch config.Env.LogLevel {
		case "debug":
			level = zap.DebugLevel
		case "warn":
			level = zap.WarnLevel
		case "error":
			level = zap.ErrorLevel
		default:
			level = zap.InfoLevel
		}

		core := zapcore.NewCore(encoder, zapcore.AddSync(logWriter), level)
		LoggerInstance = zap.New(core)
	}
}

func ZapLogger() fiber.Handler {
	var core zapcore.Core
	var level zapcore.Level

	switch config.Env.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	if config.Env.Development {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
		LoggerInstance = zap.New(core)
	} else {
		_ = os.MkdirAll("logs", 0755)
		updateLoggerIfNeeded()
	}

	return func(c *fiber.Ctx) error {
		start := time.Now()

		if !config.Env.Development {
			updateLoggerIfNeeded()
		}

		err := c.Next()
		stop := time.Since(start)

		LoggerInstance.Info("request",
			zap.String("ip", c.IP()),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", stop),
		)

		return err
	}
}
