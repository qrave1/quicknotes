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

func (nu *NoteUsecase) Create(ctx context.Context, n domain.Note) error {
	return nu.noteRepo.Add(ctx, n)
}

func (nu *NoteUsecase) Read(ctx context.Context, id int) (domain.Note, error) {
	return nu.noteRepo.GetById(ctx, id)
}

func (nu *NoteUsecase) Update(ctx context.Context, n domain.Note) error {
	return nu.noteRepo.Update(ctx, n)
}

func (nu *NoteUsecase) Delete(ctx context.Context, id int) error {
	return nu.noteRepo.Delete(ctx, id)
}
