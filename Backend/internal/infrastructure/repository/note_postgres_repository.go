package repository

import (
	"context"
	"database/sql"
	"github.com/qrave1/quicknotes/internal/domain"
)

type NotePostgresRepository struct {
	db *sql.DB
}

func NewNotePostgresRepository(db *sql.DB) *NotePostgresRepository {
	return &NotePostgresRepository{
		db: db,
	}
}

func (n *NotePostgresRepository) Add(ctx context.Context, note domain.Note) (int, error) {
	query := "INSERT INTO notes(title, body, folder_id) VALUES ($1,$2,$3) RETURNING id"
	res := n.db.QueryRowContext(ctx, query, note.Title, note.Body, note.FolderId)
	if res.Err() != nil {
		return 0, res.Err()
	}

	var id int64
	err := res.Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (n *NotePostgresRepository) GetById(ctx context.Context, id int) (domain.Note, error) {
	note := domain.Note{}
	query := "SELECT id, title, body, folder_id FROM notes WHERE id = $1"
	err := n.db.QueryRowContext(ctx, query, id).Scan(&note.Id, &note.Title, &note.Body, &note.FolderId)
	if err != nil {
		return domain.Note{}, err
	}

	return note, nil
}

func (n *NotePostgresRepository) GetAll(ctx context.Context, folderId int) ([]domain.Note, error) {
	query := `SELECT n.id, n.title, n.body, n.folder_id
				FROM folders f
				    INNER JOIN notes n ON f.id = n.folder_id
				    	WHERE f.id = $1`
	rows, err := n.db.QueryContext(ctx, query, folderId)
	if err != nil {
		return nil, err
	}

	var folderNotes []domain.Note
	for rows.Next() {
		var fn domain.Note
		err := rows.Scan(&fn.Id, &fn.Title, &fn.Body, &fn.FolderId)
		if err != nil {
			return nil, err
		}
		folderNotes = append(folderNotes, fn)
	}

	return folderNotes, nil
}

func (n *NotePostgresRepository) Update(ctx context.Context, note domain.Note) error {
	query := "UPDATE notes SET title = $1, body = $2 WHERE id = $3"
	res, err := n.db.ExecContext(ctx, query, note.Title, note.Body, note.Id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrNotAffected
	}

	return nil
}

func (n *NotePostgresRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM notes WHERE id = $1"
	res, err := n.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrNotAffected
	}

	return nil
}
