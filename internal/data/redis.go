package data

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/knadh/koanf"
)

func NewRedisClient(conf *koanf.Koanf) (*redis.Client, error) {
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
	return client, nil
}
