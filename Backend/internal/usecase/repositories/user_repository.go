package repositories

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type User interface {
	Add(ctx context.Context, u domain.User) error
	UserById(ctx context.Context, id int) (domain.User, error)
	UserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdatePass(ctx context.Context, id int, hashedPass string) error
	Delete(ctx context.Context, id int) error
}
