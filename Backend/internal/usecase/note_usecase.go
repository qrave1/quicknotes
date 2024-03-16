package usecase

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
)

type NoteService struct {
	noteRepo repositories.Note
}

func NewNoteService(nr repositories.Note) *NoteService {
	return &NoteService{noteRepo: nr}
}

func (n *NoteService) Create(ctx context.Context, note domain.Note) error {
	return n.noteRepo.Add(ctx, note)
}

func (n *NoteService) Read(ctx context.Context, id int) (domain.Note, error) {
	return n.noteRepo.GetById(ctx, id)
}

func (n *NoteService) ReadAll(ctx context.Context, folderId int) ([]domain.Note, error) {
	return n.noteRepo.GetAll(ctx, folderId)
}

func (n *NoteService) Update(ctx context.Context, note domain.Note) error {
	return n.noteRepo.Update(ctx, note)
}

func (n *NoteService) Delete(ctx context.Context, id int) error {
	return n.noteRepo.Delete(ctx, id)
}
