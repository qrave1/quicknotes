package domain

import "context"

type Folder struct {
	Id     int
	Name   string
	UserId int
}

type FolderUsecase interface {
	Create(ctx context.Context, f Folder) error
	FolderById(ctx context.Context, id int) (Folder, error)
	Folders(ctx context.Context) ([]Folder, error)
	Update(ctx context.Context, f Folder) error
	Delete(ctx context.Context, id int) error
}
