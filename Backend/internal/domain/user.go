package domain

import "context"

type User struct {
	Name     string
	Email    string
	Password string
}

type UserRepo interface {
	Add(context.Context, User) error
	GetById(context.Context, int) (User, error)
	Update(ctx context.Context, id int, u User) error
	Delete(context.Context, int) error
}

type UserUsecase interface {
	Create(ctx context.Context, u User) error
	Read(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, id int, u User) error
	Delete(ctx context.Context, id int) error
}
