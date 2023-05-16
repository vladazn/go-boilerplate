package db

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PgClient struct {
	db *bun.DB
}

func NewBunClient(config *Config) (*bun.DB, error) {
	connector := pgdriver.NewConnector(
		pgdriver.WithAddr(config.Address()),
		pgdriver.WithUser(config.Connection.User),
		pgdriver.WithPassword(config.Connection.Password),
		pgdriver.WithDatabase(config.Connection.Database),
		pgdriver.WithInsecure(true),
	)

	sqlDB := sql.OpenDB(connector)
	db := bun.NewDB(sqlDB, pgdialect.New())

	err := db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
