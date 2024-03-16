package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/qrave1/quicknotes/internal/domain"
)

type NotePostgresRepository struct {
	db   *sql.DB
	psql sq.StatementBuilderType
}

func NewNotePostgresRepository(db *sql.DB) *NotePostgresRepository {
	return &NotePostgresRepository{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (n *NotePostgresRepository) Add(ctx context.Context, note domain.Note) error {
	exec, err := n.psql.Insert("notes").
		Columns("title", "body", "folder_id").
		Values(note.Title, note.Body, note.FolderId).
		RunWith(n.db).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (n *NotePostgresRepository) GetById(ctx context.Context, id int) (domain.Note, error) {
	rows, err := n.psql.Select("*").
		From("notes").
		Where(sq.Eq{"note_id": id}).
		RunWith(n.db).
		QueryContext(ctx)
	if err != nil {
		return domain.Note{}, err
	}

	var note domain.Note
	err = rows.Scan(&note)
	if err != nil {
		return domain.Note{}, err
	}

	return note, nil
}

func (n *NotePostgresRepository) GetAll(ctx context.Context, folderId int) ([]domain.Note, error) {
	rows, err := n.psql.Select("*").From("notes").
		Join("folders ON notes.folder_id = folders.id").
		Where(sq.Eq{"folders_id": folderId}).
		RunWith(n.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var notes []domain.Note
	err = rows.Scan(&notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *NotePostgresRepository) Update(ctx context.Context, note domain.Note) error {
	query := n.psql.Update("notes")

	if note.Title != "" {
		query = query.Set("title", note.Title)
	}

	if note.Body != "" {
		query = query.Set("body", note.Body)
	}

	exec, err := query.Where(sq.Eq{"note_id": note.Id}).RunWith(n.db).ExecContext(ctx)
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (n *NotePostgresRepository) Delete(ctx context.Context, id int) error {
	exec, err := n.psql.Delete("notes").
		Where(sq.Eq{"note_id": id}).
		RunWith(n.db).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}
