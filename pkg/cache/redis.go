package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCacheService struct {
	rdb *redis.Client
}

func NewRedisCacheService(rdb *redis.Client) RedisCacheService {
	return &redisCacheService{
		rdb: rdb,
	}
}

func (cs *redisCacheService) Get(ctx context.Context, key string, dest any) error {
	data, err := cs.rdb.Get(ctx, key).Bytes()
	
	if err != nil {
		return err
	}
	
	return json.Unmarshal(data, dest)
}

func (cs *redisCacheService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		
		return err
	}

	return cs.rdb.Set(ctx, key, data, ttl).Err()
}

func (cs *redisCacheService) Exits(ctx context.Context, key string) (bool, error) {
	count, err := cs.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (cs *redisCacheService) Clear(ctx context.Context, key string) error {

	if err := cs.rdb.Del(ctx, key).Err(); err != nil && err == redis.Nil {
		return err
	}

	return nil
}

func (cs *redisCacheService) HIncrByOne(ctx context.Context, key, field string) error {
	if err := cs.rdb.HIncrBy(ctx, key, field, 1).Err(); err != nil {
		return err
	}

	return nil
}

func (cs *redisCacheService) Rename(ctx context.Context, key, newkey string) error {
	if err := cs.rdb.Rename(ctx, key, newkey).Err(); err != nil {
		return err
	}

	return nil
}

func (cs *redisCacheService) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	data, err := cs.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}, err
	}

	return data, nil
}
