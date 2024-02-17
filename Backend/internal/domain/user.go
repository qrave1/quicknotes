package domain

import "context"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type UserRepo interface {
	Add(ctx context.Context, u User) error
	GetById(ctx context.Context, id int) (User, error)
	UpdatePass(ctx context.Context, id int, hashedPass string) error
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	Create(ctx context.Context, u User) error
	Read(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, id int, hashedPass string) error
	Delete(ctx context.Context, id int) error
}
