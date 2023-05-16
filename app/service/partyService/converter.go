package partyService

import (
	"github.com/gofrs/uuid"
	"github.com/vladazn/go-boilerplate/app/storage"
)

type UpdateUsernameRequest struct {
	UserId   uuid.UUID
	Username string
}

type UpdateUserSettingsRequest struct {
	UserId   uuid.UUID
	Settings *UserSettings
}

type GetUserResponse struct {
	User *User
}

type User struct {
	UserId   uuid.UUID
	Username string
	Settings *UserSettings
}

type UserSettings struct {
	IsSoundEnabled      bool
	IsMusicEnabled      bool
	IsLeftHandedEnabled bool
}

func convertUser(userData *storage.User) *User {
	return &User{
		UserId:   userData.Id,
		Username: userData.Username,
		Settings: convertUserSettings(userData.UserSettings),
	}
}

func convertUserSettings(settingsData *storage.UserSettings) *UserSettings {
	if settingsData == nil {
		return nil
	}

	return &UserSettings{
		IsSoundEnabled:      settingsData.IsSoundEnabled,
		IsMusicEnabled:      settingsData.IsMusicEnabled,
		IsLeftHandedEnabled: settingsData.IsLeftHandedEnabled,
	}
}
