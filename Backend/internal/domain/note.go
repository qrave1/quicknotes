package domain

import "context"

type Note struct {
	Id       int
	Title    string
	Body     string
	FolderId int
}

type NoteUsecase interface {
	Create(ctx context.Context, n Note) error
	Read(ctx context.Context, id int) (Note, error)
	Update(ctx context.Context, id int, n Note) error
	Delete(ctx context.Context, id int) error
}
