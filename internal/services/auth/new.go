package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/auth/internal/config"
	"github.com/sportgroup-hq/auth/internal/models"
	"github.com/sportgroup-hq/common-lib/api"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*api.User, error)
	GetUserByEmail(ctx context.Context, email string) (*api.User, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, cur *api.CreateUserRequest) (*api.User, error)
}

type UserCredentialsRepo interface {
	GetUserCredentialsByUserID(ctx context.Context, userID uuid.UUID) (*models.UserCredential, error)
	CreateUserCredentials(ctx context.Context, credential *models.UserCredential) error
}

type KVStore interface {
	Set(ctx context.Context, key, value string, expiresIn time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Service struct {
	cfg *config.Config

	userService         UserService
	userCredentialsRepo UserCredentialsRepo

	kvStore KVStore

	googleOauthConfig *oauth2.Config
}

func NewService(cfg *config.Config, userService UserService,
	userCredentialsRepo UserCredentialsRepo, kvStore KVStore) *Service {
	return &Service{
		cfg:                 cfg,
		userService:         userService,
		userCredentialsRepo: userCredentialsRepo,
		kvStore:             kvStore,
		googleOauthConfig: &oauth2.Config{
			ClientID:     cfg.Oauth.Google.ClientID,
			ClientSecret: cfg.Oauth.Google.ClientSecret,
			RedirectURL:  cfg.Oauth.Google.RedirectURL,
			Scopes:       cfg.Oauth.Google.Scopes,
			Endpoint:     google.Endpoint,
		},
	}
}
