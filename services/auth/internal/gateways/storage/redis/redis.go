package redis_adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage"
	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	rdb *redis.Client
}

func New(url string) (*RedisAdapter, error) {
	rdbOptions, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}
	rdb := redis.NewClient(rdbOptions)
	return &RedisAdapter{rdb}, nil
}

func (r *RedisAdapter) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", storage.ErrNotFound
	}
	return val, nil
}

func (r *RedisAdapter) Set(ctx context.Context, key, value string) error {
	err := r.rdb.Set(ctx, key, value, time.Duration(0)).Err()
	return fmt.Errorf("failed to set redis value: %w", err)
}

func (r *RedisAdapter) Close(ctx context.Context) error {
	errChan := make(chan error, 1)
	go func() {
		errChan <- r.rdb.Close()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
