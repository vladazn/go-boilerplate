package grpc

import (
	"context"
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
	"github.com/vladazn/go-boilerplate/app/service/authService"
)

type AuthService interface {
	Auth(ctx context.Context, params *authService.AuthParams) (*authService.AuthResponse, error)
}

type AuthHandler struct {
	frontoffice.UnimplementedAuthServiceServer
	authService AuthService
	userService UserService
}

func NewCoreFoAuthHandler(
	authService *authService.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Auth(
	ctx context.Context,
	req *frontoffice.AuthRequest,
) (*frontoffice.AuthResponse, error) {
	err := validateAuthRequest(req)
	if err != nil {
		return nil, buildError(err)
	}

	authData := &authService.AuthParams{
		AccountType: authService.AuthAccountType(req.GetAuthType()),
		Token:       req.GetKey(),
	}

	if req.GetUsername() != nil {
		username := req.GetUsername().GetValue()
		authData.Username = &username
	}

	tokens, err := h.authService.Auth(ctx, authData)

	if err != nil {
		return nil, buildError(err)
	}

	return &frontoffice.AuthResponse{
		AccessToken: tokens.AuthToken,
	}, nil
}
