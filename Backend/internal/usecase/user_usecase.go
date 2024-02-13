package usecase

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type UserUsecase struct {
	ur domain.UserRepo
}

func (uu UserUsecase) Create(ctx context.Context, u domain.User) error {
	return uu.ur.Add(ctx, u)
}

func (uu UserUsecase) Read(ctx context.Context, id int) (domain.User, error) {
	return uu.ur.GetById(ctx, id)
}

func (uu UserUsecase) Update(ctx context.Context, id int, hashedPass string) error {
	return uu.ur.UpdatePass(ctx, id, hashedPass)
}

func (uu UserUsecase) Delete(ctx context.Context, id int) error {
	return uu.ur.Delete(ctx, id)
}
