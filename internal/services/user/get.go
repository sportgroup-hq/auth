package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/api"
)

func (s *Service) GetUserByID(ctx context.Context, userID uuid.UUID) (*api.User, error) {
	resp, err := s.api.GetUserByID(ctx, &api.GetUserByIDRequest{Id: userID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed executing remote GetUser: %w", err)
	}

	return resp.User, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*api.User, error) {
	resp, err := s.api.GetUserByEmail(ctx, &api.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, fmt.Errorf("failed executing remote GetUserByEmail: %w", err)
	}

	return resp.User, nil
}

func (s *Service) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	resp, err := s.api.UserExistsByEmail(ctx, &api.UserExistsByEmailRequest{Email: email})
	if err != nil {
		return false, fmt.Errorf("failed executing remote UserExistsByEmail: %w", err)
	}

	return resp.Exists, nil
}

func (s *Service) CreateUser(ctx context.Context, cur *api.CreateUserRequest) (*api.User, error) {
	resp, err := s.api.CreateUser(ctx, cur)
	if err != nil {
		return nil, fmt.Errorf("failed executing remote CreateUser: %w", err)
	}

	return resp.User, nil
}
