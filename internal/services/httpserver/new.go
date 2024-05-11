package httpserver

import (
	"context"

	"github.com/sportgroup-hq/auth/internal/config"
	"github.com/sportgroup-hq/auth/internal/models"
	"golang.org/x/oauth2"
)

type AuthService interface {
	GetOAuthConsentURL(ctx context.Context) string
	ExchangeProvidersCode(ctx context.Context, code string) (*oauth2.Token, error)
	LoginWithEmail(ctx context.Context, email, password string) (*oauth2.Token, error)
	Register(ctx context.Context, regDTO *models.RegisterRequest) (*oauth2.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error)
}

type HTTPServer struct {
	cfg *config.Config

	authService AuthService
}

func New(cfg *config.Config, authService AuthService) *HTTPServer {
	return &HTTPServer{
		cfg:         cfg,
		authService: authService,
	}
}
