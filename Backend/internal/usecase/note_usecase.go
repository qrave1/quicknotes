package usecase

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type NoteUsecase struct {
	nr domain.NoteRepo
}

func NewNoteUsecase(nr domain.NoteRepo) *NoteUsecase {
	return &NoteUsecase{nr: nr}
}

func (nu *NoteUsecase) Create(ctx context.Context, n domain.Note) error {
	return nu.nr.Add(ctx, n)
}

func (nu *NoteUsecase) Read(ctx context.Context, id int) (domain.Note, error) {
	return nu.nr.GetById(ctx, id)
}

func (nu *NoteUsecase) Update(ctx context.Context, id int, n domain.Note) error {
	return nu.nr.Update(ctx, id, n)
}

func (nu *NoteUsecase) Delete(ctx context.Context, id int) error {
	return nu.nr.Delete(ctx, id)
}
