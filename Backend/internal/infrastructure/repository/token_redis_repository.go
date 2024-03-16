package repository

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisNotFound struct{}

func (r RedisNotFound) Error() string {
	return ""
}

var NotFound = RedisNotFound{}

type AuthToken interface {
	Get(ctx context.Context, k string) (string, error)
	Set(ctx context.Context, k, v string) (string, error)
}

const TOKEN_TTL = 24 * time.Hour

type AuthTokenRedisRepository struct {
	redis *redis.Client
}

func NewTokenRedisRepository(c *redis.Client) *AuthTokenRedisRepository {
	return &AuthTokenRedisRepository{redis: c}
}

func (tr *AuthTokenRedisRepository) Get(ctx context.Context, k string) (string, error) {
	if res, err := tr.redis.Get(ctx, k).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return "", NotFound
		} else {
			return "", err
		}
	} else {
		return res, nil
	}
}

func (tr *AuthTokenRedisRepository) Set(ctx context.Context, k, v string) (string, error) {
	if res, err := tr.redis.Set(ctx, k, v, TOKEN_TTL).Result(); err != nil {
		return "", err
	} else {
		return res, nil
	}
}
