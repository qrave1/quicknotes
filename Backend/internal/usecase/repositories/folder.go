package repositories

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type Folder interface {
	Add(ctx context.Context, f domain.Folder) (int, error)
	GetById(ctx context.Context, id int) (domain.Folder, error)
	GetAll(ctx context.Context, userId int) ([]domain.Folder, error)
	Update(ctx context.Context, f domain.Folder) error
	Delete(ctx context.Context, id int) error
}
