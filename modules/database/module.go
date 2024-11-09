package database

import (
	"database/sql"

	"github.com/alexkazantsev/templ-api/modules/database/migrations"
	"github.com/alexkazantsev/templ-api/modules/database/storage"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(
		NewConnection,
		migrations.Inject,
		func(con *sql.DB) storage.DBTX { return con },
		storage.New,
	),
	fx.Invoke(
		Migrate,
	),
)
