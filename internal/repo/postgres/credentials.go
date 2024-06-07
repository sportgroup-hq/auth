package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/auth/internal/models"
)

func (p *Postgres) GetUserCredentialsByUserID(ctx context.Context, userID uuid.UUID) (*models.UserCredential, error) {
	var userCredential models.UserCredential

	err := p.db.NewSelect().
		Model(&userCredential).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &userCredential, nil
}

func (p *Postgres) CreateUserCredentials(ctx context.Context, userCredential *models.UserCredential) error {
	_, err := p.db.NewInsert().
		Model(userCredential).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
