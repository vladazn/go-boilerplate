package config

import (
	"github.com/vladazn/go-boilerplate/pkg/jwt"
	"time"
)

func NewJwtGeneratorConfig(config *Config) *jwt.JwtConfig {
	return &jwt.JwtConfig{
		Secret:         config.JWT.Secret,
		AuthExpiration: 3 * time.Hour,
	}
}
