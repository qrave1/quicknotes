package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/storage/postgres"
)

type FolderPostgresRepository struct {
	*postgres.Storage
	psql sq.StatementBuilderType
}

func NewFolderPostgresRepository(
	storage *postgres.Storage,
) *FolderPostgresRepository {
	return &FolderPostgresRepository{
		Storage: storage,
		psql:    sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (fr FolderPostgresRepository) Add(ctx context.Context, f domain.Folder) error {
	exec, err := fr.psql.Insert("folders").
		Columns("folder_id", "name", "user_id").
		Values(f.Id, f.Name, f.UserId).
		RunWith(fr.DB).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (fr FolderPostgresRepository) GetById(ctx context.Context, id int) (domain.Folder, error) {
	rows, err := fr.psql.Select("*").
		From("folders").
		Where(sq.Eq{"folder_id": id}).
		RunWith(fr.DB).
		Query()
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

func (fr FolderPostgresRepository) Update(ctx context.Context, id int, f domain.Folder) error {
	query := fr.psql.Update("folders")

	if f.Id != 0 {
		query = query.Set("folder_id", f.Id)
	}

	if f.Name != "" {
		query = query.Set("name", f.Name)
	}

	exec, err := query.Where(sq.Eq{"folder_id": id}).RunWith(fr.DB).Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (fr FolderPostgresRepository) Delete(ctx context.Context, id int) error {
	exec, err := fr.psql.Delete("folders").
		Where(sq.Eq{"folder_id": id}).
		RunWith(fr.DB).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}
