package config

import "github.com/vladazn/go-boilerplate/pkg/db/migration/schema"

func NewSchemaMigrationConfig() *schema.Config {
	return &schema.Config{
		MigrationPath: "db/schema",
	}
}
