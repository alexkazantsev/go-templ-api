package migrations

import "embed"

//go:embed *.sql
var Migrations embed.FS

func Inject() embed.FS {
	return Migrations
}
