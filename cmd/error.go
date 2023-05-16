package cmd

import "go.uber.org/fx"

func catchError(fxErr *error) fx.Option {
	return fx.ErrorHook(errorHandler(func(err error) {
		*fxErr = err
	}))
}

type errorHandler func(err error)

func (h errorHandler) HandleError(err error) {
	h(err)
}
