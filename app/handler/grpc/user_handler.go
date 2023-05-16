package grpc

import (
	"context"
	"github.com/gofrs/uuid"
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
	"github.com/vladazn/go-boilerplate/app/service/userService"
	"github.com/vladazn/go-boilerplate/pkg/toolkit/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	UpdateUsername(ctx context.Context, req *userService.UpdateUsernameRequest) error
	GetUserInfo(ctx context.Context, userID uuid.UUID) (*userService.GetUserResponse, error)
	SetUserSettings(ctx context.Context, params *userService.UpdateUserSettingsRequest) error
}

type UserCurrencyService interface {
	GetCurrencyCode(ctx context.Context, currencyId uuid.UUID) (string, error)
}

type UserHandler struct {
	frontoffice.UnimplementedUserServiceServer
	userService     UserService
	currencyService UserCurrencyService
}

func NewCoreFoUserHandler(
	userService *userService.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserInfo(
	ctx context.Context,
	_ *frontoffice.GetUserInfoRequest,
) (*frontoffice.GetUserInfoResponse, error) {
	userId := auth.UserIdFromContext(ctx)
	if userId == uuid.Nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	userInfo, err := h.userService.GetUserInfo(ctx, userId)
	if err != nil {
		return nil, buildError(err)
	}

	return &frontoffice.GetUserInfoResponse{
		Info: convertUserInfo(userInfo.User),
	}, nil
}

func (h *UserHandler) SetUserSettings(
	ctx context.Context,
	req *frontoffice.SetUserSettingsRequest,
) (*frontoffice.SetUserSettingsResponse, error) {
	userId := auth.UserIdFromContext(ctx)
	if userId == uuid.Nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if err := h.userService.SetUserSettings(ctx, &userService.UpdateUserSettingsRequest{
		UserId: userId,
		Settings: &userService.UserSettings{
			IsSoundEnabled:      req.Settings.IsSoundEnabled,
			IsMusicEnabled:      req.Settings.IsMusicEnabled,
			IsLeftHandedEnabled: req.Settings.IsLeftHandedEnabled,
		},
	}); err != nil {
		return nil, buildError(err)
	}

	return &frontoffice.SetUserSettingsResponse{}, nil
}

func (h *UserHandler) UpdateUsername(
	ctx context.Context,
	req *frontoffice.SetUsernameRequest,
) (*frontoffice.SetUsernameResponse, error) {
	userId := auth.UserIdFromContext(ctx)
	if userId == uuid.Nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err := validateUpdateUsernameRequest(req)
	if err != nil {
		return nil, buildError(err)
	}

	if err := h.userService.UpdateUsername(ctx, &userService.UpdateUsernameRequest{
		UserId:   userId,
		Username: req.Username,
	}); err != nil {
		return nil, buildError(err)
	}

	return &frontoffice.SetUsernameResponse{}, nil
}
