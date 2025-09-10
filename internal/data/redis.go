package data

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/internal/system"
)

type RedisClient struct {
	Client *redis.Client
	conf   *koanf.Koanf
	ctx    context.Context

	rcf *payloads.RedisConfig
}

func NewRedisClient(conf *koanf.Koanf, rcf *payloads.RedisConfig, app *system.AppContext) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rcf.Addr,
		Password: rcf.Password,
		DB:       rcf.DB,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	rc := &RedisClient{
		conf:   conf,
		Client: client,
		ctx:    app.Ctx,
		rcf:    rcf,
	}

	return rc, nil
}

func (rc *RedisClient) CachedExpire(ctx context.Context, key string, val any) error {
	return rc.Client.Set(ctx, key, val, 0).Err()
}
