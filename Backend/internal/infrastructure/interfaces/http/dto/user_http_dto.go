package dto

import "github.com/qrave1/quicknotes/internal/domain"

type UserRequest interface {
	Name() string
	Email() string
	Password() string
}

type SignInRequest struct {
	email    string `form:"email" validate:"required,email"`
	password string `form:"password" validate:"required, min=8, max=30"`
}

func (s *SignInRequest) Name() string {
	return ""
}

func (s *SignInRequest) Email() string {
	return s.email
}

func (s *SignInRequest) Password() string {
	return s.password
}

type SignUpRequest struct {
	name     string `form:"name" validate:"required"`
	email    string `form:"email" validate:"required,email"`
	password string `form:"password" validate:"required, min=8, max=30"`
}

func (s *SignUpRequest) Name() string {
	return s.name
}

func (s *SignUpRequest) Email() string {
	return s.email
}

func (s *SignUpRequest) Password() string {
	return s.password
}

func UserFromDTO(r UserRequest) domain.User {
	return domain.User{
		Id:       0,
		Name:     r.Name(),
		Email:    r.Email(),
		Password: r.Password(),
	}
}
