package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/auth/internal/models"
	"github.com/sportgroup-hq/common-lib/api"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) LoginWithEmail(ctx context.Context, email, password string) (*oauth2.Token, error) {
	user, err := s.userService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	userID, err := uuid.Parse(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}

	userCredentials, err := s.userCredentialsRepo.GetUserCredentialsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user credentials by email: %w", err)
	}

	if userCredentials == nil {
		return nil, errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userCredentials.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("wrong password")
	}

	// Create token
	token, err := s.createBearerToken(ctx, userID.String(), s.cfg.JWT.AccessExpiresIn, s.cfg.JWT.RefreshExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("failed to create bearer token: %w", err)
	}

	return token, nil
}

func (s *Service) Register(ctx context.Context, regDTO *models.RegisterRequest) (*oauth2.Token, error) {
	user, err := s.userService.GetUserByEmail(ctx, regDTO.Email)
	if err != nil {
		if err, ok := status.FromError(err); !(ok && err.Code() == codes.NotFound) {
			return nil, fmt.Errorf("failed to get user by email: %s", err)
		}
	}

	if user != nil {
		userID, err := uuid.Parse(user.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to parse user id: %w", err)
		}

		userCredentials, err := s.userCredentialsRepo.GetUserCredentialsByUserID(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("failed to get user credentials by user id: %w", err)
		}

		if userCredentials != nil {
			return nil, errors.New("user already exists")
		}
	} else {
		cur := &api.CreateUserRequest{
			Email:     regDTO.Email,
			FirstName: regDTO.FirstName,
			LastName:  regDTO.LastName,
		}

		if user, err = s.userService.CreateUser(ctx, cur); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	userID, err := uuid.Parse(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}

	userCredentials := &models.UserCredential{
		UserID:       userID,
		PasswordHash: string(hashedPassword),
	}

	if err = s.userCredentialsRepo.CreateUserCredentials(ctx, userCredentials); err != nil {
		return nil, fmt.Errorf("failed to create user credentials: %w", err)
	}

	// Create token
	token, err := s.createBearerToken(ctx, userID.String(), s.cfg.JWT.AccessExpiresIn, s.cfg.JWT.RefreshExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("failed to create bearer token: %w", err)
	}

	return token, nil
}
