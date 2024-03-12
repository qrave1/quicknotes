package dto

import "github.com/qrave1/quicknotes/internal/domain"

type UserRequest interface {
	GetName() string
	GetEmail() string
	GetPassword() string
}

type SignInRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required, min=8, max=30"`
}

func (s *SignInRequest) GetName() string {
	return ""
}

func (s *SignInRequest) GetEmail() string {
	return s.Email
}

func (s *SignInRequest) GetPassword() string {
	return s.Password
}

type SignUpRequest struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required, min=8, max=30"`
}

func (s *SignUpRequest) GetName() string {
	return s.Name
}

func (s *SignUpRequest) GetEmail() string {
	return s.Email
}

func (s *SignUpRequest) GetPassword() string {
	return s.Password
}

func UserFromDTO(r UserRequest) domain.User {
	return domain.User{
		Id:       0,
		Name:     r.GetName(),
		Email:    r.GetEmail(),
		Password: r.GetPassword(),
	}
}
