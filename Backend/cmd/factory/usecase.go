package factory

import (
	"github.com/google/wire"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/infrastructure/repository"
	"github.com/qrave1/quicknotes/internal/usecase"
	"github.com/qrave1/quicknotes/internal/usecase/auth"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
)

var usecaseSet = wire.NewSet(
	wire.Bind(new(domain.NoteUsecase), new(*usecase.NoteService)),
	provideNoteUsecase,

	wire.Bind(new(domain.FolderUsecase), new(*usecase.FolderService)),
	provideFolderUsecase,

	wire.Bind(new(domain.UserUsecase), new(*usecase.UserService)),
	provideUserUsecase,

	wire.Bind(new(auth.Auth), new(*auth.AuthService)),
	provideAuthUsecase,
)

func provideNoteUsecase(nr repositories.Note) *usecase.NoteService {
	return usecase.NewNoteService(nr)
}

func provideFolderUsecase(fr repositories.Folder) *usecase.FolderService {
	return usecase.NewFolderService(fr)
}

func provideUserUsecase(
	userRepo repositories.User,
	tokenRepo repository.AuthToken,
	auth auth.Auth,
	log logrus.Logger,
) *usecase.UserService {
	return usecase.NewUserService(userRepo, tokenRepo, auth, log)
}

func provideAuthUsecase(cfg *config.Config, log logrus.Logger) *auth.AuthService {
	return auth.NewAuthService(cfg.Server.Secret, log)
}
