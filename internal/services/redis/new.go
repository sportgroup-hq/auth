package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sportgroup-hq/auth/internal/config"
)

type Service struct {
	cli *redis.Client
}

func New(cfg *config.Config) (*Service, error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &Service{
		cli: redisCli,
	}, nil
}
