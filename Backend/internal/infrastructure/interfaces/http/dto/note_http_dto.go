package dto

import "github.com/qrave1/quicknotes/internal/domain"

type NoteRequest interface {
	GetId() int
	GetTitle() string
	GetBody() string
	GetFolderId() int
}

type CreateNoteRequest struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	FolderId int    `path:"folder_id" validate:"required"`
}

func (c CreateNoteRequest) GetId() int {
	return 0
}

func (c CreateNoteRequest) GetID() int {
	return 0
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
	Id       int    `path:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	FolderId int    `path:"folder_id" validate:"required"`
}

func (c UpdateNoteRequest) GetId() int {
	return 0
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

type DefaultNoteRequest struct {
	NoteId   int `path:"id"`
	FolderId int `path:"folder_id" validate:"required"`
}

func (d DefaultNoteRequest) GetId() int {
	return d.NoteId
}

func (d DefaultNoteRequest) GetTitle() string {
	return ""
}

func (d DefaultNoteRequest) GetBody() string {
	return ""
}

func (d DefaultNoteRequest) GetFolderId() int {
	return d.FolderId
}

// NoteFromDTO Используется только для запросов из этого пакета
// Потому что я в рот

//func NoteFromDTO(r NoteRequest) domain.Note {
//	return domain.Note{
//		Id:       0,
//		Title:    r.GetTitle(),
//		Body:     r.GetBody(),
//		FolderId: r.GetFolderId(),
//	}
//}

func NoteFromDTO[T NoteRequest](req T) domain.Note {
	return domain.Note{
		Id:       req.GetId(),
		Title:    req.GetTitle(),
		Body:     req.GetBody(),
		FolderId: req.GetFolderId(),
	}
}
