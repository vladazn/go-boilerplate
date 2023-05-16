package schema

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"io/fs"
)

type Migrator struct {
	config     *Config
	db         *bun.DB
	migrations *migrate.Migrations
}

type MigrationsFileSystem fs.FS

func NewMigrator(
	config *Config,
	db *bun.DB,
	migrationFS MigrationsFileSystem,
) (*Migrator, error) {
	migrations := migrate.NewMigrations(
		migrate.WithMigrationsDirectory(config.MigrationPath),
	)
	err := migrations.Discover(migrationFS)
	if err != nil {
		return nil, err
	}

	migrator := &Migrator{
		config:     config,
		db:         db,
		migrations: migrations,
	}

	return migrator, nil
}

func (m *Migrator) Migrate(ctx context.Context) error {
	migrator := migrate.NewMigrator(
		m.db,
		m.migrations,
	)

	err := migrator.Init(ctx)
	if err != nil {
		return err
	}

	_, err = migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	return nil
}
