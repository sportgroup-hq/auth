package bootstrap

import (
	"github.com/sportgroup-hq/auth/internal/config"
	"github.com/sportgroup-hq/auth/internal/services/auth"
	userGRPC "github.com/sportgroup-hq/auth/internal/services/user"
	"github.com/sportgroup-hq/common-lib/api"
)

type Dependencies struct {
	Config config.Config

	//HTTPServer *httpserver.HTTPServer
	AuthService *auth.Service

	UserService *userGRPC.Service

	GRPCApiClient api.ApiClient
}

func NewDependencies(config config.Config, authService *auth.Service,
	userService *userGRPC.Service, grpcApiClient api.ApiClient) *Dependencies {
	return &Dependencies{
		Config:        config,
		AuthService:   authService,
		UserService:   userService,
		GRPCApiClient: grpcApiClient,
	}
}
