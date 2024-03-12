package dto

import "github.com/qrave1/quicknotes/internal/domain"

type FolderRequest interface {
	Id() int
	Name() string
	UserId() int
}

type CreateFolderRequest struct {
	name string `json:"name" validate:"required"`
}

func (c CreateFolderRequest) Id() int {
	return 0
}

func (c CreateFolderRequest) Name() string {
	return c.name
}

func (c CreateFolderRequest) UserId() int {
	return 0
}

func FolderFromDTO(r FolderRequest) domain.Folder {
	return domain.Folder{
		Id:     r.Id(),
		Name:   r.Name(),
		UserId: r.UserId(),
	}
}
