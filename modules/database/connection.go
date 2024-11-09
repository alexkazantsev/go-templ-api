package database

import (
	"database/sql"

	"github.com/alexkazantsev/go-templ-api/modules/config"
	"github.com/avast/retry-go"
	"go.uber.org/zap"
)

func NewConnection(cfg config.DatabaseConfig, logger *zap.Logger) (*sql.DB, error) {
	var (
		db *sql.DB
	)

	var err = retry.Do(func() error {
		var (
			dbErr error
		)

		if db, dbErr = sql.Open("postgres", cfg.GetDSN()); dbErr != nil {
			return dbErr
		}

		return db.Ping()
	},
		retry.Attempts(5),
		retry.OnRetry(func(n uint, err error) {
			logger.Warn("could not connect to database, retrying", zap.Uint("attempt", n), zap.Error(err))
		}),
	)

	if err != nil {
		return nil, err
	}

	logger.Info("connected to database")

	return db, nil
}
