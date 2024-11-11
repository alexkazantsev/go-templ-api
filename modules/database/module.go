package database

import (
	"database/sql"

	"github.com/alexkazantsev/go-templ-api/modules/database/storage"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(
		NewConnection,
		func(con *sql.DB) storage.DBTX { return con },
		storage.New,
	),
)
