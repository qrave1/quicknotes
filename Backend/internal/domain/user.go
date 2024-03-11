package domain

import "context"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type UserUsecase interface {
	SignUp(ctx context.Context, user User) error
	SignIn(ctx context.Context, request User) (string, error)
	Read(ctx context.Context, id int) (User, error)
	Update(ctx context.Context, id int, hashedPass string) error
	Delete(ctx context.Context, id int) error
}
