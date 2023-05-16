package schema

import (
	"embed"
	"github.com/vladazn/go-boilerplate/pkg/db/migration/schema"
	"go.uber.org/fx"
)

var (
	//go:embed *.sql
	migrations embed.FS
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			func() schema.MigrationsFileSystem {
				return migrations
			},
		),
	)
}
