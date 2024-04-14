package usecase

import (
	"context"
	"fmt"
	"github.com/qrave1/logwrap"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/interface/errors"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
)

type FolderService struct {
	folderRepo repositories.Folder
	log        logwrap.Logger
}

func NewFolderService(fr repositories.Folder, log logwrap.Logger) *FolderService {
	return &FolderService{folderRepo: fr, log: log}
}

func (f *FolderService) Create(ctx context.Context, folder domain.Folder) (int, error) {
	folder.UserId = auth.UserIdFromCtx(ctx)

	return f.folderRepo.Add(ctx, folder)
}

func (f *FolderService) FolderById(ctx context.Context, id int) (domain.Folder, error) {
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

func (f *FolderService) Folders(ctx context.Context) ([]domain.Folder, error) {
	currentUserId := auth.UserIdFromCtx(ctx)

	folders, err := f.folderRepo.GetAll(ctx, currentUserId)
	if err != nil {
		return nil, fmt.Errorf("error get folder from repo. %w", err)
	}

	for _, folder := range folders {
		if folder.UserId != currentUserId {
			return nil, errors.ForbiddenError
		}
	}

	return folders, nil
}

func (f *FolderService) Update(ctx context.Context, folder domain.Folder) error {
	return f.folderRepo.Update(ctx, folder)
}

func (f *FolderService) Delete(ctx context.Context, id int) error {
	return f.folderRepo.Delete(ctx, id)
}
