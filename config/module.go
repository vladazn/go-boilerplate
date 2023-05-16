package config

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewConfig,
			NewStorageConfig,
			NewSchemaMigrationConfig,
			NewGrpcServerConfigs,
			NewGatewayServerConfigs,
			NewJwtGeneratorConfig,
		),
	)
}
