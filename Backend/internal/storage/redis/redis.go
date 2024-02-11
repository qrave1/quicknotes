package redis

import (
	"context"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	redis *redis.Client
}

func MustLoad(c *config.Config) *Storage {
	opts, err := redis.ParseURL(c.CacheConfig.DSN)
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(opts)

	err = cli.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return &Storage{redis: cli}
}
