package partyService

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
	"github.com/vladazn/go-boilerplate/app/storage"
)

type PartyRepository interface {
	InsertOne(ctx context.Context, user *storage.Party) error
}

type PartyService struct {
	db              *bun.DB
	partyRepository PartyRepository
}

func NewPartyService(
	connection *bun.DB,
	partyRepository *storage.PartyRepository,
) *PartyService {
	return &PartyService{
		db:              connection,
		partyRepository: partyRepository,
	}
}

func (h *PartyService) CreateParty(ctx context.Context, partyName string) error {
	party := &storage.Party{
		Id:        uuid.Must(uuid.NewV4()),
		PartyName: partyName,
	}

	err := h.partyRepository.InsertOne(ctx, party)
	if err != nil {
		return err
	}

	return nil
}
