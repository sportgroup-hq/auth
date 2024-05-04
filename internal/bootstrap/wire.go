//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/sportgroup-hq/auth/internal/config"
	"github.com/sportgroup-hq/auth/internal/services/auth"
	userGRPC "github.com/sportgroup-hq/auth/internal/services/user"
	"github.com/sportgroup-hq/common-lib/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Up() (*Dependencies, error) {
	wire.Build(
		config.New,

		//httpserver.New,
		newGRPCService,

		userGRPC.New,

		auth.NewService,
		wire.Bind(new(auth.UserService), new(*userGRPC.Service)),

		NewDependencies,
	)
	return &Dependencies{}, nil
}

func newGRPCService(cfg config.Config) (api.ApiClient, error) {
	clientConn, err := grpc.NewClient(cfg.Services.API.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed creating new grpc client: %w", err)
	}

	return api.NewApiClient(clientConn), nil
}
