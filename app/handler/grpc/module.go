package grpc

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewCoreFoAuthHandler,
			NewCoreFoUserHandler,
			NewPartyHandler,
		),
	)
}
