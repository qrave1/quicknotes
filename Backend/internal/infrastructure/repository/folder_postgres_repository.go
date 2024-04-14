package repository

import (
	"context"
	"database/sql"
	"github.com/qrave1/quicknotes/internal/domain"
)

type FolderPostgresRepository struct {
	db *sql.DB
}

func NewFolderPostgresRepository(db *sql.DB) *FolderPostgresRepository {
	return &FolderPostgresRepository{
		db: db,
	}
}

func (f *FolderPostgresRepository) Add(ctx context.Context, folder domain.Folder) (int, error) {
	query := "INSERT INTO folders(name, user_id) VALUES ($1,$2) RETURNING id"
	res := f.db.QueryRowContext(ctx, query, folder.Name, folder.UserId)
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

func (f *FolderPostgresRepository) GetById(ctx context.Context, id int) (domain.Folder, error) {
	folder := domain.Folder{}
	query := "SELECT id, name, user_id FROM folders WHERE id = $1"
	err := f.db.QueryRowContext(ctx, query, id).Scan(&folder.Id, &folder.Name, &folder.UserId)
	if err != nil {
		return domain.Folder{}, err
	}

	return folder, nil
}

func (f *FolderPostgresRepository) GetAll(ctx context.Context, userId int) ([]domain.Folder, error) {
	query := `SELECT f.id, f.name, f.user_id
				FROM users u
				    INNER JOIN folders f ON f.user_id = u.id
				    	WHERE f.user_id = $1
				    		ORDER BY f.id`
	rows, err := f.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var userFolders []domain.Folder
	for rows.Next() {
		var uf domain.Folder
		err := rows.Scan(&uf.Id, &uf.Name, &uf.UserId)
		if err != nil {
			return nil, err
		}
		userFolders = append(userFolders, uf)
	}

	return userFolders, nil
}

func (f *FolderPostgresRepository) Update(ctx context.Context, folder domain.Folder) error {
	query := "UPDATE folders SET name = $1 WHERE id = $2"
	res, err := f.db.ExecContext(ctx, query, folder.Name, folder.Id)
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

func (f *FolderPostgresRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM folders WHERE id = $1"
	res, err := f.db.ExecContext(ctx, query, id)
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
