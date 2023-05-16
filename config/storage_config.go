package config

import "github.com/vladazn/go-boilerplate/pkg/db"

func NewStorageConfig(config *Config) *db.Config {
	return &db.Config{
		Connection: db.Connection{
			Host:     config.Pg.Host,
			Port:     config.Pg.Port,
			User:     config.Pg.User,
			Password: config.Pg.Pass,
			Database: config.Pg.Db,
		},
	}
}
