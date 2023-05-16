package gateway

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewServer,
		),
		fx.Invoke(
			runServer,
		),
	)
}
