package bootstrap

import (
	"github.com/sportgroup-hq/auth/internal/config"
	"github.com/sportgroup-hq/auth/internal/controller/httpserver"
	"github.com/sportgroup-hq/auth/internal/services/auth"
	"github.com/sportgroup-hq/auth/internal/services/postgres"
	"github.com/sportgroup-hq/auth/internal/services/redis"
	userGRPC "github.com/sportgroup-hq/auth/internal/services/user"
	"github.com/sportgroup-hq/common-lib/api"
)

type Dependencies struct {
	Config *config.Config

	HTTPServer *httpserver.Server

	AuthService *auth.Service

	UserService *userGRPC.Service

	GRPCApiClient api.ApiClient

	Postgres *postgres.Postgres
	Redis    *redis.Service
}

func NewDependencies(config *config.Config, httpServer *httpserver.Server, authService *auth.Service,
	userService *userGRPC.Service, grpcApiClient api.ApiClient, postgres *postgres.Postgres,
	redis *redis.Service) *Dependencies {
	return &Dependencies{
		Config:        config,
		HTTPServer:    httpServer,
		AuthService:   authService,
		UserService:   userService,
		GRPCApiClient: grpcApiClient,
		Postgres:      postgres,
		Redis:         redis,
	}
}
