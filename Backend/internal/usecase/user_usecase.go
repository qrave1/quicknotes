package usecase

import (
	"context"
	"github.com/qrave1/quicknotes/internal/domain"
)

type UserUsecase struct {
	ur domain.UserRepo
}

func (uu UserUsecase) Create(ctx context.Context, u domain.User) error {

}

func (uu UserUsecase) Read(ctx context.Context, id int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uu UserUsecase) Update(ctx context.Context, id int, u domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (uu UserUsecase) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
