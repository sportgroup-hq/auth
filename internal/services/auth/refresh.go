package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sportgroup-hq/auth/internal/models"
	"golang.org/x/oauth2"
)

type Claims struct {
	jwt.RegisteredClaims
	Type string `json:"typ"`
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	var claims Claims

	parsedToken, err := jwt.ParseWithClaims(refreshToken, &claims, s.jwtSecretFunc)
	if err != nil || !parsedToken.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired: %w", err)
		}

		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims.Type != "refresh" {
		return nil, errors.New("you must provide a refresh token")
	}

	userIDKey := models.RefreshTokenUserIDKey(claims.ID)

	userID, err := s.kvStore.Get(ctx, userIDKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get user id from kv-store: %w", err)
	}

	if userID == "" {
		return nil, errors.New("invalid token")
	}

	if err = s.kvStore.Delete(ctx, userIDKey); err != nil {
		return nil, fmt.Errorf("failed to delete refresh token from kv-store: %w", err)
	}

	token, err := s.createBearerToken(ctx, userID, s.cfg.JWT.AccessExpiresIn, s.cfg.JWT.RefreshExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("failed to create bearer token: %w", err)
	}

	return token, nil
}

func (s *Service) jwtSecretFunc(_ *jwt.Token) (interface{}, error) {
	return []byte(s.cfg.JWT.Secret), nil
}
