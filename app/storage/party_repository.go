package storage

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/vladazn/go-boilerplate/pkg/db"
)

type PartyRepository struct {
	db *bun.DB
}

func NewPartyRepository(db *bun.DB) *PartyRepository {
	return &PartyRepository{db: db}
}

func (r *PartyRepository) InsertOne(ctx context.Context, user *Party) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return db.NormalizeError(err)
	}

	return nil
}
