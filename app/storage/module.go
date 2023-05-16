package storage

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewUserRepository,
			//user additional data
			NewUserSettingsRepository,
		),
	)
}
