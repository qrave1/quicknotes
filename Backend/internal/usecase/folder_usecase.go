package usecase

import (
	"context"
	"fmt"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/interface/errors"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
)

type FolderUsecase struct {
	folderRepo repositories.Folder
	log        logrus.Logger
}

func NewFolderUsecase(fr repositories.Folder) *FolderUsecase {
	return &FolderUsecase{folderRepo: fr}
}

func (f *FolderUsecase) Create(ctx context.Context, folder domain.Folder) error {
	folder.UserId = auth.UserIdFromCtx(ctx)

	return f.folderRepo.Add(ctx, folder)
}

func (f *FolderUsecase) Read(ctx context.Context, id int) (domain.Folder, error) {
	currentUserId := auth.UserIdFromCtx(ctx)

	folder, err := f.folderRepo.GetById(ctx, id)
	if err != nil {
		return domain.Folder{}, fmt.Errorf("error get folder from repo. %w", err)
	}

	if folder.UserId != currentUserId {
		return domain.Folder{}, errors.ForbiddenError
	}

	return folder, nil
}

func (f *FolderUsecase) Update(ctx context.Context, folder domain.Folder) error {
	return f.folderRepo.Update(ctx, folder)
}

func (f *FolderUsecase) Delete(ctx context.Context, id int) error {
	return f.folderRepo.Delete(ctx, id)
}
