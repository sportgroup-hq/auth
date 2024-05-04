package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/api"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*api.User, error)
}

type Service struct {
	userService UserService
}

func NewService(userService UserService) *Service {
	return &Service{
		userService: userService,
	}
}

func (s *Service) test() error {
	s.userService.GetUserByID(context.Background(), uuid.New())
	return nil
}
