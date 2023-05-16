package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func runServer(
	lc fx.Lifecycle,
	server *Server,
	errs chan error,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logrus.Info(ctx, "Running gRPC server...")

			go func() {
				errs <- server.Start()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Info(ctx, "Stopping gRPC server...")

			server.Stop()

			return nil
		},
	})
}
