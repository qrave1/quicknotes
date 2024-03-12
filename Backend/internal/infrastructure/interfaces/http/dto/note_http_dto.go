package dto

import "github.com/qrave1/quicknotes/internal/domain"

type NoteRequest interface {
	GetTitle() string
	GetBody() string
	GetFolderId() int
}

type CreateNoteRequest struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	FolderId int    `path:"folder_id" validate:"required"`
}

func (c CreateNoteRequest) GetTitle() string {
	return c.Title
}

func (c CreateNoteRequest) GetBody() string {
	return c.Body
}

func (c CreateNoteRequest) GetFolderId() int {
	return c.FolderId
}

type UpdateNoteRequest struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	FolderId int    `path:"folder_id" validate:"required"`
}

func (c UpdateNoteRequest) GetTitle() string {
	return c.Title
}

func (c UpdateNoteRequest) GetBody() string {
	return c.Body
}

func (c UpdateNoteRequest) GetFolderId() int {
	return c.FolderId
}

func NoteFromDTO(r NoteRequest) domain.Note {
	return domain.Note{
		Id:       0,
		Title:    r.GetTitle(),
		Body:     r.GetBody(),
		FolderId: r.GetFolderId(),
	}
}
