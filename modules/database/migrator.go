package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/alexkazantsev/go-templ-api/modules/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file includes file driver
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/hashicorp/go-multierror"
	"go.uber.org/zap"
)

func Migrate(db *sql.DB, params config.DatabaseConfig, logger *zap.Logger, migrations embed.FS) error {
	src, err := iofs.New(migrations, ".")
	if err != nil {
		return fmt.Errorf("iofs new migration source: %w", err)
	}

	if err := up(logger, db, src, params.Name, ""); err != nil {
		return fmt.Errorf("database migrate: %w", err)
	}

	return nil
}

func up(
	l *zap.Logger, db *sql.DB, src source.Driver, dbName string, dsn string,
) error {
	var (
		m   *migrate.Migrate
		err error
	)

	if dsn == "" {
		driver, drvErr := postgres.WithInstance(db, &postgres.Config{
			DatabaseName: dbName,
		})
		if drvErr != nil {
			return fmt.Errorf("postgres with instance: %w", drvErr)
		}

		m, err = migrate.NewWithInstance("iofs", src, dbName, driver)
		if err != nil {
			return fmt.Errorf("migrate new with instance: %w", err)
		}
	} else {
		m, err = migrate.NewWithSourceInstance("iofs", src, dsn)
		if err != nil {
			return fmt.Errorf("migrate new with source instance: %w", err)
		}
	}

	startVersion, _, err := m.Version()
	l.Info("start version", zap.Uint("version", startVersion))

	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return fmt.Errorf("migration get version: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return multierror.Append(fmt.Errorf("migrate up: %w", err), rollbackToStartMigration(m, startVersion))
	}

	v, d, err := m.Version()
	if err != nil {
		return fmt.Errorf("migration get version: %w", err)
	}

	l.Info("migration done", zap.Uint("version", v), zap.Bool("dirty", d))

	return nil
}

func rollbackToStartMigration(m *migrate.Migrate, startVersion uint) error {
	currentVersion, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("version: %w", err)
	}

	// remove dirty value from schema
	if dirty {
		currentVersion -= 1
		if err := m.Force(int(currentVersion)); err != nil {
			return fmt.Errorf("force is fail: %w", err)
		}
	}

	// current version rollback to start version
	if currentVersion != startVersion {
		if err := m.Steps(int(startVersion - currentVersion)); err != nil {
			return multierror.Append(fmt.Errorf("rollback up: %w", err), clearDirtyMigration(m, 1))
		}
	}

	return nil
}

func clearDirtyMigration(m *migrate.Migrate, step int) error {
	v, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("version: %w", err)
	}

	if dirty {
		if err := m.Force(int(v) + step); err != nil {
			return fmt.Errorf("force is fail: %w", err)
		}
	}

	return nil
}
