package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/qrave1/quicknotes/internal/domain"
)

type FolderPostgresRepository struct {
	db   *sql.DB
	psql sq.StatementBuilderType
}

func NewFolderPostgresRepository(db *sql.DB) *FolderPostgresRepository {
	return &FolderPostgresRepository{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (f *FolderPostgresRepository) Add(ctx context.Context, folder domain.Folder) error {
	exec, err := f.psql.Insert("folders").
		Columns("name", "user_id").
		Values(folder.Name, folder.UserId).
		RunWith(f.db).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (f *FolderPostgresRepository) GetById(ctx context.Context, id int) (domain.Folder, error) {
	rows, err := f.psql.Select("*").
		From("folders").
		Where(sq.Eq{"folder_id": id}).
		RunWith(f.db).
		QueryContext(ctx)
	if err != nil {
		return domain.Folder{}, err
	}

	var folder domain.Folder
	err = rows.Scan(&folder)
	if err != nil {
		return domain.Folder{}, err
	}

	return folder, nil
}

func (f *FolderPostgresRepository) GetAll(ctx context.Context, userId int) ([]domain.Folder, error) {
	rows, err := f.psql.Select("*").From("folders").
		Join("users u ON folders.user_id = u.id").
		Where(sq.Eq{"user_id": userId}).
		RunWith(f.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var folders []domain.Folder
	err = rows.Scan(&folders)
	if err != nil {
		return nil, err
	}

	return folders, nil
}

func (f *FolderPostgresRepository) Update(ctx context.Context, folder domain.Folder) error {
	query := f.psql.Update("folders")

	if folder.Name != "" {
		query = query.Set("name", folder.Name)
	}

	exec, err := query.Where(sq.Eq{"folder_id": folder.Id}).RunWith(f.db).ExecContext(ctx)
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (f *FolderPostgresRepository) Delete(ctx context.Context, id int) error {
	exec, err := f.psql.Delete("folders").
		Where(sq.Eq{"folder_id": id}).
		RunWith(f.db).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}
