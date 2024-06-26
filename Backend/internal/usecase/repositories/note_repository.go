package repositories

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type Note interface {
	Add(ctx context.Context, n domain.Note) (int, error)
	GetById(ctx context.Context, id int) (domain.Note, error)
	GetAll(ctx context.Context, folderId int) ([]domain.Note, error)
	Update(ctx context.Context, n domain.Note) error
	Delete(ctx context.Context, id int) error
}
