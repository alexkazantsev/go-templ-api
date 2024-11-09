package config

import (
	"net"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.uber.org/zap"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

func (e Environment) IsProduction() bool {
	return e == Production
}

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Fatal LogLevel = "fatal"
)

func (l LogLevel) ToZapLevel() zap.AtomicLevel {
	switch l {
	case Debug:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case Info:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case Warn:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case Error:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case Fatal:
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

type AppConfig struct {
	Name        string      `env:"APP_NAME"`
	Environment Environment `env:"APP_ENV"`
	LogLevel    LogLevel    `env:"LOG_LEVEL"`
	Port        string      `env:"APP_PORT"`
}

func (c AppConfig) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Environment, validation.Required, validation.In(Development, Production)),
		validation.Field(&c.LogLevel, validation.Required, validation.In(Debug, Info, Warn, Error, Fatal)),
		validation.Field(&c.Port, validation.Required, is.Port),
	)
}

func (c AppConfig) GetAddr() string {
	return net.JoinHostPort("", c.Port)
}
