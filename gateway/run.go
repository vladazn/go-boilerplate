package gateway

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func runServer(
	lc fx.Lifecycle,
	server *GatewayServer,
	errs chan error,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logrus.Info(ctx, "Running gateway server...")

			go func() {
				errs <- server.Start()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Info(ctx, "Stopping gateway server...")

			return server.Stop()
		},
	})
}
