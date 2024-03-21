package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/qrave1/quicknotes/internal/domain"
)

var (
	ErrNotAffected = errors.New("no rows affected")
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		db: db,
	}
}

func (ur *UserPostgresRepository) Add(ctx context.Context, user domain.User) error {
	query := "INSERT INTO users(username, email, password) VALUES ($1,$2,$3)"
	return ur.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Err()
}

func (ur *UserPostgresRepository) UserById(ctx context.Context, id int) (domain.User, error) {
	u := domain.User{}
	query := "SELECT id, username, email, password FROM users WHERE id = $1"
	err := ur.db.QueryRowContext(ctx, query, id).Scan(&u.Id, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (ur *UserPostgresRepository) UserByEmail(ctx context.Context, email string) (domain.User, error) {
	u := domain.User{}
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	err := ur.db.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (ur *UserPostgresRepository) UpdatePass(ctx context.Context, id int, hashedPass string) error {
	query := "UPDATE users SET password = $1 WHERE id = $2"
	res, err := ur.db.ExecContext(ctx, query, hashedPass, id)
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

func (ur *UserPostgresRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	res, err := ur.db.ExecContext(ctx, query, id)
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
