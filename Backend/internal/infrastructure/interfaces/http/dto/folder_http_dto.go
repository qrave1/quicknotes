package dto

import "github.com/qrave1/quicknotes/internal/domain"

type FolderRequest interface {
	GetId() int
	GetName() string
	GetUserId() int
}

type CreateFolderRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateFolderRequest) GetId() int {
	return 0
}

func (c CreateFolderRequest) GetName() string {
	return c.Name
}

func (c CreateFolderRequest) GetUserId() int {
	return 0
}

type UpdateFolderRequest struct {
	Name string `json:"name" validate:"required"`
	Id   int    `path:"id" validate:"required"`
}

func (u UpdateFolderRequest) GetId() int {
	return u.Id
}

func (u UpdateFolderRequest) GetName() string {
	return u.Name
}

func (u UpdateFolderRequest) GetUserId() int {
	return 0
}

func FolderFromDTO(r FolderRequest) domain.Folder {
	return domain.Folder{
		Id:     r.GetId(),
		Name:   r.GetName(),
		UserId: r.GetUserId(),
	}
}
