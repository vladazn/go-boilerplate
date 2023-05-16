package authService

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
	"github.com/vladazn/go-boilerplate/app/storage"
	"github.com/vladazn/go-boilerplate/pkg/jwt"
	"sync"
	"time"
)

type UserRepository interface {
	FindOneByTokenAndPlatformForUpdate(ctx context.Context, token string, platform storage.UserPlatform) (*storage.User, error)
	InsertOne(ctx context.Context, user *storage.User) error
	UpdateOne(ctx context.Context, user *storage.User, columns []string) error
}

type UserSettingsRepository interface {
	InsertOne(ctx context.Context, user *storage.UserSettings) error
}

type JwtGenerator interface {
	GenerateAuthToken(userID uuid.UUID) string
	GenerateRefreshToken(userID uuid.UUID) string
	ParseAuthToken(tokenString string) (*jwt.AuthTokenData, error)
	ParseRefreshToken(tokenString string) (*jwt.RefreshTokenData, error)
}

type AuthService struct {
	db                     *bun.DB
	userRepository         UserRepository
	userSettingsRepository UserSettingsRepository
	jwtGenerator           JwtGenerator
	lock                   *sync.Mutex
	usernamePool           []string
}

func NewAuthService(
	db *bun.DB,
	userRepository *storage.UserRepository,
	userSettingsRepository *storage.UserSettingsRepository,
	jwtGenerator *jwt.JwtGenerator,
) *AuthService {
	return &AuthService{
		db:                     db,
		userRepository:         userRepository,
		userSettingsRepository: userSettingsRepository,
		jwtGenerator:           jwtGenerator,
		lock:                   &sync.Mutex{},
	}
}

func (s *AuthService) Auth(ctx context.Context, params *AuthParams) (*AuthResponse, error) {
	var user *storage.User

	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, _ bun.Tx) error {
		var err error
		user, err = s.userRepository.FindOneByTokenAndPlatformForUpdate(
			ctx,
			params.Token,
			storage.UserPlatform(params.AccountType),
		)
		if err != nil {
			return err
		}

		if user == nil {
			user, err = s.createNewUser(ctx, params)
			if err != nil {
				return err
			}
		} else {
			user.LastLoginAt = time.Now()
			err = s.userRepository.UpdateOne(ctx, user, []string{"last_login_at"})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	authToken := s.jwtGenerator.GenerateAuthToken(user.Id)

	return &AuthResponse{
		AuthToken: authToken,
	}, nil
}

func (s *AuthService) createNewUser(ctx context.Context, params *AuthParams) (*storage.User, error) {
	user := &storage.User{
		Id:          uuid.Must(uuid.NewV4()),
		Token:       params.Token,
		Platform:    storage.UserPlatform(params.AccountType),
		LastLoginAt: time.Now(),
	}

	if params.Username == nil {
		user.Username = "GUEST"
	} else {
		user.Username = *params.Username
	}

	err := s.userRepository.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	err = s.registerUserSettings(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) registerUserSettings(ctx context.Context, userId uuid.UUID) error {
	err := s.userSettingsRepository.InsertOne(ctx, &storage.UserSettings{
		Id:                  uuid.Must(uuid.NewV4()),
		UserId:              userId,
		IsSoundEnabled:      true,
		IsMusicEnabled:      true,
		IsLeftHandedEnabled: false,
	})
	if err != nil {
		return err
	}

	return nil
}
