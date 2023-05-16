package storage

import (
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
)

type Party struct {
	bun.BaseModel `bun:"boilerplate.party"`

	Id        uuid.UUID `bun:"type:uuid,pk"`
	PartyName string    `bun:"type:text"`
}
