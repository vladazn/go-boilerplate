package grpc

import (
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
	"github.com/vladazn/go-boilerplate/app/service/userService"
)

func convertUserInfo(info *userService.User) *frontoffice.UserInfo {
	if info == nil {
		return nil
	}

	return &frontoffice.UserInfo{
		UserId:   info.UserId.String(),
		Username: info.Username,
		Settings: convertUserSettings(info.Settings),
	}
}

func convertUserSettings(settings *userService.UserSettings) *frontoffice.UserSettings {
	if settings == nil {
		return nil
	}

	return &frontoffice.UserSettings{
		IsSoundEnabled:      settings.IsSoundEnabled,
		IsMusicEnabled:      settings.IsMusicEnabled,
		IsLeftHandedEnabled: settings.IsLeftHandedEnabled,
	}
}
