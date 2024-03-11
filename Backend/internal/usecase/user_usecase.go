package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/repository"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo  repositories.User
	tokenRepo repository.AuthTokenRepo
	auth      auth.Auth
	log       logrus.Logger
}

func (u *UserUsecase) SignUp(ctx context.Context, user domain.User) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("error generate hashed password. %v", err)
		return err
	}
	user.Password = string(passHash)

	return u.userRepo.Add(ctx, user)
}

func (u *UserUsecase) SignIn(ctx context.Context, request domain.User) (string, error) {
	user, err := u.userRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if user.Email != request.Email {
		return "", fmt.Errorf("error email does not match")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(request.Password), []byte(user.Password)); err != nil {
		u.log.Infof("invalid credentials. user_id=%d", request.Id)
		return "", err
	}

	token, err := u.tokenRepo.Get(ctx, user.Email)
	if err != nil {
		if errors.Is(err, repository.NotFound) {
			token, err = u.auth.Generate(user.Id)
			if err != nil {
				return "", err
			}
			_, err = u.tokenRepo.Set(ctx, user.Email, token)
		} else {
			return "", nil
		}
	}

	return token, nil
}

func (u *UserUsecase) Read(ctx context.Context, id int) (domain.User, error) {
	return u.userRepo.GetById(ctx, id)
}

func (u *UserUsecase) Update(ctx context.Context, id int, pass string) error {
	return u.userRepo.UpdatePass(ctx, id, pass)
}

func (u *UserUsecase) Delete(ctx context.Context, id int) error {
	return u.userRepo.Delete(ctx, id)
}
