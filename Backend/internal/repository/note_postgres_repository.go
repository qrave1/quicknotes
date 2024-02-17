package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/storage/postgres"
)

type NotePostgresRepository struct {
	*postgres.Storage
	psql sq.StatementBuilderType
}

func NewNotePostgresRepository(storage *postgres.Storage) *NotePostgresRepository {
	return &NotePostgresRepository{
		Storage: storage,
		psql:    sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (nr *NotePostgresRepository) Add(ctx context.Context, n domain.Note) error {
	exec, err := nr.psql.Insert("notes").
		Columns("note_id", "title", "body", "folder_id").
		Values(n.Id, n.Title, n.Body, n.FolderId).
		RunWith(nr.DB).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (nr *NotePostgresRepository) GetById(ctx context.Context, id int) (domain.Note, error) {
	rows, err := nr.psql.Select("*").
		From("notes").
		Where(sq.Eq{"note_id": id}).
		RunWith(nr.DB).
		Query()
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

func (nr *NotePostgresRepository) Update(ctx context.Context, id int, n domain.Note) error {
	query := nr.psql.Update("notes")

	if n.Title != "" {
		query = query.Set("title", n.Title)
	}

	if n.Body != "" {
		query = query.Set("body", n.Body)
	}

	exec, err := query.Where(sq.Eq{"note_id": id}).RunWith(nr.DB).Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (nr *NotePostgresRepository) Delete(ctx context.Context, id int) error {
	exec, err := nr.psql.Delete("notes").
		Where(sq.Eq{"note_id": id}).
		RunWith(nr.DB).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}
