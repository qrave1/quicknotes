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
	Create()
	Read()
	Update()
	Delete()
}
