package logger

import (
	"fmt"

	"github.com/alexkazantsev/templ-api/modules/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg config.AppConfig) (*zap.Logger, error) {
	var (
		err error
		lg  *zap.Logger
	)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	c := zap.Config{
		Level:             cfg.LogLevel.ToZapLevel(),
		Development:       !cfg.Environment.IsProduction(),
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableCaller:     true,
		DisableStacktrace: true,
	}

	if lg, err = c.Build(); err != nil {
		return nil, fmt.Errorf("could not create logger: %w", err)
	}

	defer func(lg *zap.Logger) {
		_ = lg.Sync()
	}(lg)

	lg.Info("logger created")

	return lg, nil
}
