package storage

import (
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
)

type UserSettings struct {
	bun.BaseModel `bun:"boilerplate.user_settings"`

	Id                  uuid.UUID `bun:"type:uuid,pk"`
	UserId              uuid.UUID `bun:"type:uuid"`
	IsSoundEnabled      bool      `bun:"type:boolean"`
	IsMusicEnabled      bool      `bun:"type:boolean"`
	IsLeftHandedEnabled bool      `bun:"type:boolean"`
}
