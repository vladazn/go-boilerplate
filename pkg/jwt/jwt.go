package jwt

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

type AuthTokenData struct {
	UserId    uuid.UUID
	ExpiresAt time.Time
}

func (a *AuthTokenData) IsValid() bool {
	return a.ExpiresAt.After(time.Now()) && a.UserId != uuid.Nil
}

type RefreshTokenData struct {
	UserId    uuid.UUID
	ExpiresAt time.Time
}

type JwtGenerator struct {
	config *JwtConfig
}

func NewJwtGenerator(config *JwtConfig) *JwtGenerator {
	return &JwtGenerator{
		config: config,
	}
}

func (g *JwtGenerator) GenerateAuthToken(userID uuid.UUID) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     userID,
		"expires_at": time.Now().Add(g.config.AuthExpiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(g.config.Secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (g *JwtGenerator) GenerateRefreshToken(userID uuid.UUID) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     userID.String(),
		"expires_at": time.Now().Add(g.config.RefreshExpiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(g.config.Secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (g *JwtGenerator) ParseAuthToken(tokenString string) (*AuthTokenData, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	params := AuthTokenData{}

	userId, ok := claims["userId"]
	if !ok {
		return nil, errors.New("missing user id")
	}

	params.UserId, err = uuid.FromString(userId.(string))
	if err != nil {
		return nil, errors.New("missing user id")
	}

	expires, ok := claims["expires_at"]
	if !ok {
		return nil, errors.New("invalid timestamp")
	}

	expiresAt, ok := expires.(float64)
	if !ok {
		return nil, errors.New("invalid timestamp")
	}

	params.ExpiresAt = time.Unix(int64(expiresAt), 0)

	return &params, nil
}

func (g *JwtGenerator) ParseRefreshToken(tokenString string) (*RefreshTokenData, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	params := RefreshTokenData{}

	userId, ok := claims["userId"]
	if !ok {
		return nil, errors.New("missing user id")
	}

	params.UserId, err = uuid.FromString(userId.(string))
	if err != nil {
		return nil, errors.New("missing user id")
	}

	expires, ok := claims["expires_at"]
	if !ok {
		return nil, errors.New("invalid timestamp")
	}

	expiresAt, ok := expires.(float64)
	if !ok {
		return nil, errors.New("invalid timestamp")
	}

	params.ExpiresAt = time.Unix(int64(expiresAt), 0)

	return &params, nil
}
