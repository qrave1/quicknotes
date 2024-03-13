package usecase

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
)

type NoteUsecase struct {
	noteRepo repositories.Note
}

func NewNoteUsecase(nr repositories.Note) *NoteUsecase {
	return &NoteUsecase{noteRepo: nr}
}

func (n *NoteUsecase) Create(ctx context.Context, note domain.Note) error {
	return n.noteRepo.Add(ctx, note)
}

func (n *NoteUsecase) Read(ctx context.Context, id int) (domain.Note, error) {
	return n.noteRepo.GetById(ctx, id)
}

func (n *NoteUsecase) ReadAll(ctx context.Context, folderId int) ([]domain.Note, error) {
	return n.noteRepo.GetAll(ctx, folderId)
}

func (n *NoteUsecase) Update(ctx context.Context, note domain.Note) error {
	return n.noteRepo.Update(ctx, note)
}

func (n *NoteUsecase) Delete(ctx context.Context, id int) error {
	return n.noteRepo.Delete(ctx, id)
}
