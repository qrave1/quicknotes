package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/qrave1/logwrap"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/repository"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  repositories.User
	tokenRepo repository.AuthToken
	auth      auth.Auth
	log       logwrap.Logger
}

func NewUserService(
	userRepo repositories.User,
	tokenRepo repository.AuthToken,
	auth auth.Auth,
	log logwrap.Logger,
) *UserService {
	return &UserService{userRepo: userRepo, tokenRepo: tokenRepo, auth: auth, log: log}
}

func (u *UserService) SignUp(ctx context.Context, user domain.User) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("error generate hashed password. %v", err)
		return err
	}
	user.Password = string(passHash)

	return u.userRepo.Add(ctx, user)
}

func (u *UserService) SignIn(ctx context.Context, request domain.User) (string, error) {
	user, err := u.userRepo.UserByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if user.Email != request.Email {
		return "", fmt.Errorf("error email does not match")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
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

func (u *UserService) Read(ctx context.Context, id int) (domain.User, error) {
	return u.userRepo.UserById(ctx, id)
}

func (u *UserService) Update(ctx context.Context, id int, pass string) error {
	return u.userRepo.UpdatePass(ctx, id, pass)
}

func (u *UserService) Delete(ctx context.Context, id int) error {
	return u.userRepo.Delete(ctx, id)
}
