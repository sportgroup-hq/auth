package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/api"
)

func (s *Service) GetUserByID(ctx context.Context, userID uuid.UUID) (*api.User, error) {
	resp, err := s.api.GetUser(ctx, &api.GetUserRequest{Id: userID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed executing remote GetUser: %w", err)
	}

	return resp.User, nil
}
