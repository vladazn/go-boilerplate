package storage

import (
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"boilerplate.user"`

	Id          uuid.UUID    `bun:"type:uuid,pk"`
	Token       string       `bun:"type:text"`
	Platform    UserPlatform `bun:"type:smallint"`
	Username    string       `bun:"type:text"`
	LastLoginAt time.Time    `bun:"type:timestamp"`

	UserSettings *UserSettings `bun:"rel:has-one"`
}

type UserPlatform int

const (
	UserPlatformGoogle UserPlatform = 1
	UserPlatformApple  UserPlatform = 2
)
