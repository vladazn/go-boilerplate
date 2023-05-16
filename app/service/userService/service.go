package userService

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
	"github.com/vladazn/go-boilerplate/app/storage"
)

type DatabaseConnection interface {
	Transaction(ctx context.Context, callback func(ctx context.Context) error) error
}

type UserRepository interface {
	UpdateOne(ctx context.Context, user *storage.User, columns []string) error
	FindOneByID(ctx context.Context, id uuid.UUID) (*storage.User, error)
	FineOneWithRelationsByID(ctx context.Context, id uuid.UUID) (*storage.User, error)
}

type UserSettingsRepository interface {
	FindOneByUserID(ctx context.Context, userId uuid.UUID) (*storage.UserSettings, error)
	UpdateOne(ctx context.Context, userSettings *storage.UserSettings, columns []string) error
}

type UserService struct {
	db                     *bun.DB
	userRepository         UserRepository
	userSettingsRepository UserSettingsRepository
}

func NewUserService(
	connection *bun.DB,
	userRepository *storage.UserRepository,
	userSettingsRepository *storage.UserSettingsRepository,
) *UserService {
	return &UserService{
		db:                     connection,
		userRepository:         userRepository,
		userSettingsRepository: userSettingsRepository,
	}
}

func (s *UserService) GetUserInfo(ctx context.Context, userID uuid.UUID) (*GetUserResponse, error) {
	userData, err := s.userRepository.FineOneWithRelationsByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, newNotFoundError(userID)
	}

	return &GetUserResponse{
		User: convertUser(userData),
	}, nil
}

func (s *UserService) SetUserSettings(ctx context.Context, req *UpdateUserSettingsRequest) error {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, _ bun.Tx) error {
		user, err := s.userRepository.FineOneWithRelationsByID(ctx, req.UserId)
		if err != nil {
			return err
		}
		if user == nil {
			return newNotFoundError(req.UserId)
		}

		userSettings := user.UserSettings
		userSettings.IsSoundEnabled = req.Settings.IsSoundEnabled
		userSettings.IsLeftHandedEnabled = req.Settings.IsLeftHandedEnabled
		userSettings.IsMusicEnabled = req.Settings.IsMusicEnabled

		err = s.userSettingsRepository.UpdateOne(ctx,
			userSettings,
			[]string{"is_sound_enabled", "is_music_enabled", "is_left_handed_enabled"},
		)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserSettings(ctx context.Context, userID uuid.UUID) (*UserSettings, error) {
	userSettingsData, err := s.userSettingsRepository.FindOneByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userSettingsData == nil {
		return nil, newNotFoundError(userID)
	}

	return convertUserSettings(userSettingsData), err
}

func (s *UserService) UpdateUsername(ctx context.Context, req *UpdateUsernameRequest) error {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, _ bun.Tx) error {
		user, err := s.userRepository.FindOneByID(ctx, req.UserId)
		if err != nil {
			return err
		}
		if user == nil {
			return newNotFoundError(req.UserId)
		}

		user.Username = req.Username

		err = s.userRepository.UpdateOne(ctx, user, []string{"username"})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
