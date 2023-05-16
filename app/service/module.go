package service

import (
	"github.com/vladazn/go-boilerplate/app/service/authService"
	"github.com/vladazn/go-boilerplate/app/service/partyService"
	"github.com/vladazn/go-boilerplate/app/service/userService"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			authService.NewAuthService,
			userService.NewUserService,
			partyService.NewPartyService,
		),
	)
}
