package cache

import (
	"context"
	"time"
)

type RedisCacheService interface {
	Get(ctx context.Context, key string, dest any) error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Exits(ctx context.Context, key string) (bool, error)
	Clear(ctx context.Context, key string) error
	HIncrByOne(ctx context.Context, key, field string) error
	Rename(ctx context.Context, key, newkey string) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}