package storage

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
	"github.com/vladazn/go-boilerplate/pkg/db"
)

type UserSettingsRepository struct {
	db *bun.DB
}

func NewUserSettingsRepository(db *bun.DB) *UserSettingsRepository {
	return &UserSettingsRepository{db: db}
}

func (r *UserSettingsRepository) InsertOne(ctx context.Context, userSettings *UserSettings) error {
	_, err := r.db.NewInsert().Model(userSettings).Exec(ctx)
	if err != nil {
		return db.NormalizeError(err)
	}

	return nil
}

func (r *UserSettingsRepository) FindOneByUserID(ctx context.Context, userID uuid.UUID) (*UserSettings, error) {
	userSettings := &UserSettings{}

	err := r.db.
		NewSelect().
		Model(userSettings).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, db.NormalizeError(err)
	}

	return userSettings, nil
}

func (r *UserSettingsRepository) UpdateOne(ctx context.Context, userSettings *UserSettings, columns []string) error {
	_, err := r.db.NewUpdate().Model(userSettings).Column(columns...).WherePK().Exec(ctx)
	if err != nil {
		return db.NormalizeError(err)
	}

	return nil
}
