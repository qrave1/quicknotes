package repository

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/storage/redis"
	"time"
)

const TOKEN_TTL = 24 * time.Hour

type AuthTokenRedisRepository struct {
	*redis.Storage
}

func NewTokenRedisRepository(storage *redis.Storage) *AuthTokenRedisRepository {
	return &AuthTokenRedisRepository{Storage: storage}
}

func (tr AuthTokenRedisRepository) Get(ctx context.Context, k string) (domain.AuthToken, error) {
	if res, err := tr.Storage.Redis.Get(ctx, k).Result(); err != nil {
		return domain.AuthToken{}, err
	} else {
		return domain.AuthToken{
			Key:   k,
			Value: res,
		}, nil
	}
}

func (tr AuthTokenRedisRepository) Set(ctx context.Context, k, v string) (domain.AuthToken, error) {
	if res, err := tr.Storage.Redis.Set(ctx, k, v, TOKEN_TTL).Result(); err != nil {
		return domain.AuthToken{}, err
	} else {
		return domain.AuthToken{
			Key:   k,
			Value: res,
		}, nil
	}
}
