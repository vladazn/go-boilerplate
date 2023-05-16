package jwt

import "time"

type JwtConfig struct {
	Secret            string
	AuthExpiration    time.Duration
	RefreshExpiration time.Duration
}
