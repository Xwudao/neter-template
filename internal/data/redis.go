package data

import (
	"context"
	"time"

	"github.com/Xwudao/neter-template/internal/system"
	"github.com/go-redis/redis/v8"
	"github.com/knadh/koanf"
)

type RedisClient struct {
	Client *redis.Client
	conf   *koanf.Koanf
	ctx    context.Context
}

func NewRedisClient(conf *koanf.Koanf, app *system.AppContext) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.String("redis.addr"),
		Password: conf.String("redis.password"),
		DB:       conf.Int("redis.db"),
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	rc := &RedisClient{
		conf:   conf,
		Client: client,
		ctx:    app.Ctx,
	}

	return rc, nil
}

func (rc *RedisClient) CachedExpire(key string, expir time.Duration) (existed bool) {
	ctx := rc.ctx
	_, err := rc.Client.Get(ctx, key).Result()
	if err == nil {
		return true
	}
	rc.Client.SetNX(ctx, key, "1", expir)

	return false
}
