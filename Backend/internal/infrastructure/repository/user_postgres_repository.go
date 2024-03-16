package repository

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/qrave1/quicknotes/internal/domain"
)

var (
	ErrNotAffected = errors.New("no rows affected")
)

type UserPostgresRepository struct {
	db   *sql.DB
	psql sq.StatementBuilderType
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ur *UserPostgresRepository) Add(ctx context.Context, user domain.User) error {
	exec, err := ur.psql.Insert("users").
		Columns("name", "password", "email").
		Values(user.Name, user.Password, user.Email).
		RunWith(ur.db).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (ur *UserPostgresRepository) GetById(ctx context.Context, id int) (domain.User, error) {
	rows, err := ur.psql.Select("*").
		From("users").
		Where(sq.Eq{"user_id": id}).
		RunWith(ur.db).
		Query()
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = rows.Scan(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *UserPostgresRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	rows, err := ur.psql.Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		RunWith(ur.db).
		Query()
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = rows.Scan(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *UserPostgresRepository) UpdatePass(ctx context.Context, id int, hashedPass string) error {
	exec, err := ur.psql.Update("users").
		Set("password", hashedPass).
		Where(sq.Eq{"user_id": id}).
		RunWith(ur.db).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}

func (ur *UserPostgresRepository) Delete(ctx context.Context, id int) error {
	exec, err := ur.psql.Delete("users").
		Where(sq.Eq{"user_id": id}).
		RunWith(ur.db).
		Exec()
	if err != nil {
		return err
	}

	if i, _ := exec.RowsAffected(); i == 0 {
		return ErrNotAffected
	}

	return nil
}
