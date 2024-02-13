package domain

import (
	"context"
)

type AuthToken struct {
	Key   string
	Value string
}

type AuthTokenRepo interface {
	Get(ctx context.Context, k string) (AuthToken, error)
	Set(ctx context.Context, k, v string) (AuthToken, error)
}
