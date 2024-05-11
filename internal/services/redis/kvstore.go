package redis

import (
	"context"
	"time"
)

func (s *Service) Set(ctx context.Context, key, value string, expiresIn time.Duration) error {
	return s.cli.Set(ctx, key, value, expiresIn).Err()
}

func (s *Service) Get(ctx context.Context, key string) (string, error) {
	return s.cli.Get(ctx, key).Result()
}

// Delete deletes a key from the redis store
func (s *Service) Delete(ctx context.Context, key string) error {
	return s.cli.Del(ctx, key).Err()
}
