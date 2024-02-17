package repository

import (
	"context"
	"github.com/qrave1/quicknotes/internal/storage/redis"
	"time"
)

type AuthToken struct {
	Key   string
	Value string
}

type AuthTokenRepo interface {
	Get(ctx context.Context, k string) (AuthToken, error)
	Set(ctx context.Context, k, v string) (AuthToken, error)
}

const TOKEN_TTL = 24 * time.Hour

type AuthTokenRedisRepository struct {
	*redis.Storage
}

func NewTokenRedisRepository(storage *redis.Storage) *AuthTokenRedisRepository {
	return &AuthTokenRedisRepository{Storage: storage}
}

func (tr *AuthTokenRedisRepository) Get(ctx context.Context, k string) (AuthToken, error) {
	if res, err := tr.Storage.Redis.Get(ctx, k).Result(); err != nil {
		return AuthToken{}, err
	} else {
		return AuthToken{
			Key:   k,
			Value: res,
		}, nil
	}
}

func (tr *AuthTokenRedisRepository) Set(ctx context.Context, k, v string) (AuthToken, error) {
	if res, err := tr.Storage.Redis.Set(ctx, k, v, TOKEN_TTL).Result(); err != nil {
		return AuthToken{}, err
	} else {
		return AuthToken{
			Key:   k,
			Value: res,
		}, nil
	}
}
