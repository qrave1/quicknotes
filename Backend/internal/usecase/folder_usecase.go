package usecase

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type FolderUsecase struct {
	fr domain.FolderRepo
}

func NewFolderUsecase(fr domain.FolderRepo) *FolderUsecase {
	return &FolderUsecase{fr: fr}
}

func (fu FolderUsecase) Create(ctx context.Context, f domain.Folder) error {
	return fu.fr.Add(ctx, f)
}

func (fu FolderUsecase) Read(ctx context.Context, id int) (domain.Folder, error) {
	return fu.fr.GetById(ctx, id)
}

func (fu FolderUsecase) Update(ctx context.Context, id int, f domain.Folder) error {
	return fu.fr.Update(ctx, id, f)
}

func (fu FolderUsecase) Delete(ctx context.Context, id int) error {
	return fu.fr.Delete(ctx, id)
}
