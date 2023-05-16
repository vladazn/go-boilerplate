package partyService

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type NotFoundError struct {
	id uuid.UUID
}

func newNotFoundError(id uuid.UUID) *NotFoundError {
	return &NotFoundError{id: id}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("user with %v id not found", e.id)
}
