package cmd

import (
	"github.com/vladazn/go-boilerplate/app/handler/grpc"
	"github.com/vladazn/go-boilerplate/app/server"
	"github.com/vladazn/go-boilerplate/app/service"
	"github.com/vladazn/go-boilerplate/app/storage"
	"github.com/vladazn/go-boilerplate/config"
	"github.com/vladazn/go-boilerplate/db/schema"
	"github.com/vladazn/go-boilerplate/gateway"
	"github.com/vladazn/go-boilerplate/pkg/db"
	dbschema "github.com/vladazn/go-boilerplate/pkg/db/migration/schema"
	"github.com/vladazn/go-boilerplate/pkg/jwt"
	"go.uber.org/fx"
)

func coreModule() fx.Option {
	return fx.Options(
		config.Module(),
		jwt.Module(),
		db.Module(),
		storage.Module(),
		service.Module(),
		grpc.Module(),
		server.Module(),
	)
}

func schemaMigratorModule() fx.Option {
	return fx.Options(
		config.Module(),
		db.Module(),
		schema.Module(),
		dbschema.Module(),
	)
}

func gatewayModule() fx.Option {
	return fx.Options(
		config.Module(),
		jwt.Module(),
		gateway.Module(),
	)
}
