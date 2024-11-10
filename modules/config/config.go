package config

import (
	"github.com/caarlos0/env/v6"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	Database    DatabaseConfig
	Application AppConfig
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Database),
		validation.Field(&c.Application),
	)
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	if err := cfg.Validate(); err != nil {
		return cfg, err
	}

	return Config{
		Database:    cfg.Database,
		Application: cfg.Application,
	}, nil
}
