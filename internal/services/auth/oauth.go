package auth

import (
	"context"
	"fmt"

	"github.com/sportgroup-hq/common-lib/api"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func (s *Service) GetOAuthConsentURL(ctx context.Context) string {
	return s.googleOauthConfig.AuthCodeURL(randString(64))
}

func (s *Service) ExchangeProvidersCode(ctx context.Context, code string) (*oauth2.Token, error) {
	googleToken, err := s.googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	// Parse ID token
	idTokenPayload, err := idtoken.ParsePayload(googleToken.Extra("id_token").(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ID token: %w", err)
	}

	email := idTokenPayload.Claims["email"].(string)

	exists, err := s.userService.UserExistsByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}

	var user *api.User

	if exists {
		user, err = s.userService.GetUserByEmail(ctx, email)
		if err != nil {
			return nil, fmt.Errorf("failed to get user by email: %w", err)
		}
	} else {
		picture := idTokenPayload.Claims["picture"].(string)

		user, err = s.userService.CreateUser(ctx, &api.CreateUserRequest{
			Email:     email,
			FirstName: idTokenPayload.Claims["given_name"].(string),
			LastName:  idTokenPayload.Claims["family_name"].(string),
			Picture:   &picture,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	token, err := s.createBearerToken(ctx, user.Id, s.cfg.JWT.AccessExpiresIn, s.cfg.JWT.RefreshExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("failed to own oauth token: %w", err)
	}

	return token, nil
}
