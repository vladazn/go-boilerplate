package storage

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
	"github.com/vladazn/go-boilerplate/pkg/db"
)

type UserRelations string

const (
	UserRelationUserSettings UserRelations = "UserSettings"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertOne(ctx context.Context, user *User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return db.NormalizeError(err)
	}

	return nil
}

func (r *UserRepository) UpdateOne(ctx context.Context, user *User, columns []string) error {
	_, err := r.db.NewUpdate().Model(user).Column(columns...).WherePK().Exec(ctx)
	if err != nil {
		return db.NormalizeError(err)
	}

	return nil
}

func (r *UserRepository) FineOneByID(ctx context.Context, id uuid.UUID) (*User, error) {
	user := &User{}
	err := r.db.
		NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, db.NormalizeError(err)
	}

	return user, nil
}

func (r *UserRepository) FindOneByTokenAndPlatformForUpdate(
	ctx context.Context,
	token string,
	platform UserPlatform,
) (*User, error) {
	user := &User{}
	err := r.db.
		NewSelect().
		Model(user).
		Where("token = ?", token).
		Where("platform = ?", int(platform)).
		For("update").
		Scan(ctx)
	if err != nil {
		return nil, db.NormalizeError(err)
	}

	return user, nil
}

func (r *UserRepository) FineOneWithRelationsByID(ctx context.Context, id uuid.UUID) (*User, error) {
	user := &User{}
	err := r.db.
		NewSelect().
		Model(user).
		Relation(string(UserRelationUserSettings)).
		Where("\"user\".\"id\" = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, db.NormalizeError(err)
	}

	return user, nil
}

func (r *UserRepository) FindOneByID(ctx context.Context, id uuid.UUID) (*User, error) {
	user := &User{}
	err := r.db.
		NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, db.NormalizeError(err)
	}

	return user, nil
}
