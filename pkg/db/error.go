package db

import (
	"database/sql"
	"github.com/pkg/errors"
)

func NormalizeError(err error) error {
	if errors.Cause(err) == sql.ErrNoRows {
		return nil
	}

	return err
}
