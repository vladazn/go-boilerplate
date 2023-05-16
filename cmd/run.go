package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func runApp(modules fx.Option) (rErr error) {
	var (
		err error
	)

	logrus.SetLevel(logrus.DebugLevel)

	runtimeErrs := make(chan error)

	app := fx.New(
		fx.Supply(runtimeErrs),
		modules,
		catchError(&err),
	)

	if err != nil {
		return err
	}

	err = app.Start(context.Background())
	if err != nil {
		return errors.WithStack(err)
	}

	select {
	case <-app.Done():
	case err = <-runtimeErrs:
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}
