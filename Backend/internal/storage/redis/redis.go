package redis

import (
	"context"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/redis/go-redis/v9"
)

func MustLoad(c *config.Config) *redis.Client {
	opts, err := redis.ParseURL(c.CacheConfig.DSN)
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(opts)

	err = cli.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return cli
}
