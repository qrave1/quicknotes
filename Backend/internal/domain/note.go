package domain

import "context"

type Note struct {
	Id       int
	Title    string
	Body     string
	FolderId int
}

type NoteUsecase interface {
	Create(ctx context.Context, n Note) (int, error)
	Read(ctx context.Context, id int) (Note, error)
	ReadAll(ctx context.Context, folderId int) ([]Note, error)
	Update(ctx context.Context, n Note) error
	Delete(ctx context.Context, id int) error
}
