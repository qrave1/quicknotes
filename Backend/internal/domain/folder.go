package domain

import "context"

type Folder struct {
	Id     int
	Name   string
	UserId int
}

type FolderRepo interface {
	Add(context.Context, Folder) error
	GetById(context.Context, int) (Folder, error)
	Update(ctx context.Context, id int, f Folder) error
	Delete(context.Context, int) error
}

type FolderUsecase interface {
	Create(ctx context.Context, f Folder) error
	Read(ctx context.Context, id int) (Folder, error)
	Update(ctx context.Context, id int, f Folder) error
	Delete(ctx context.Context, id int) error
}
