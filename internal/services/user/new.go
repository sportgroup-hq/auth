package user

import "github.com/sportgroup-hq/common-lib/api"

type Service struct {
	api api.ApiClient
}

func New(api api.ApiClient) *Service {
	return &Service{api}
}
