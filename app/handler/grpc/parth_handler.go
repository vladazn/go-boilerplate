package grpc

import (
	"context"
	frontoffice "github.com/vladazn/go-boilerplate/api/client"
	"github.com/vladazn/go-boilerplate/app/service/partyService"
)

type PartyService interface {
	CreateParty(ctx context.Context, partyName string) error
}

type PartyHandler struct {
	frontoffice.UnimplementedPartyServiceServer
	partyService PartyService
}

func NewPartyHandler(
	partyService *partyService.PartyService,
) *PartyHandler {
	return &PartyHandler{
		partyService: partyService,
	}
}

func (h *PartyHandler) CreateParty(
	ctx context.Context,
	req *frontoffice.CreatePartyRequest,
) (*frontoffice.CreatePartyResponse, error) {
	err := h.partyService.CreateParty(ctx, req.PartyName)
	if err != nil {
		return nil, buildError(err)
	}

	return &frontoffice.CreatePartyResponse{}, nil
}
