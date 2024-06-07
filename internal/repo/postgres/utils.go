package postgres

import (
	"database/sql"
	"errors"
)

func (p *Postgres) err(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
