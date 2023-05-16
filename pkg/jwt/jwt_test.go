package jwt

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJwtGenerator(t *testing.T) {
	generator := NewJwtGenerator(&JwtConfig{
		Secret:            "somesecret",
		AuthExpiration:    2 * time.Minute,
		RefreshExpiration: 5 * time.Minute,
	})

	userId := uuid.Must(uuid.FromString("c5b8f5a0-5c3a-4b5e-8c1d-8c5c8c5c8c5c"))

	token := generator.GenerateRefreshToken(userId)

	data, err := generator.ParseAuthToken(token)
	require.NoError(t, err)
	fmt.Println(data)

	require.Equal(t, uuid.Must(uuid.FromString("c5b8f5a0-5c3a-4b5e-8c1d-8c5c8c5c8c5c")), data.UserId)
	require.True(t, data.IsValid())
}

func TestJwtGeneratorInvalid(t *testing.T) {
	generator := NewJwtGenerator(&JwtConfig{
		Secret:         "somesecret",
		AuthExpiration: -2 * time.Minute,
	})

	userId := uuid.Must(uuid.FromString("c5b8f5a0-5c3a-4b5e-8c1d-8c5c8c5c8c5c"))

	token := generator.GenerateAuthToken(userId)

	data, err := generator.ParseAuthToken(token)
	require.NoError(t, err)

	require.Equal(t, uuid.Must(uuid.FromString("c5b8f5a0-5c3a-4b5e-8c1d-8c5c8c5c8c5c")), data.UserId)
	require.False(t, data.IsValid())
}

func TestJwtToken(t *testing.T) {
	generator := NewJwtGenerator(&JwtConfig{
		Secret:         "secretpass",
		AuthExpiration: -2 * time.Minute,
	})

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzX2F0IjoxNjc5ODYyODQ4LCJ1c2VySWQiOiJmNjcyZjhiYy05MjU0LTRlMGQtYjY3NC01Njg4Yzg2YWQ1ODEifQ.fYW2_V95zxIzavwY0iW-4wPEc77ao69PGrCqLSzWpLI"
	data, err := generator.ParseAuthToken(token)
	require.NoError(t, err)
	_ = data
}
