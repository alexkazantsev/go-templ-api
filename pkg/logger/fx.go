package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var WithZapLogger = fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: logger}
})
