package factory

import (
	"github.com/google/wire"
	"github.com/qrave1/logger-wrapper/logrus"
	"github.com/qrave1/quicknotes/internal/domain"
	"github.com/qrave1/quicknotes/internal/interface/controller"
)

var controllerSet = wire.NewSet(
	wire.Bind(new(controller.Note), new(*controller.NoteController)),
	provideNoteController,

	wire.Bind(new(controller.Folder), new(*controller.FolderController)),
	provideFolderController,

	wire.Bind(new(controller.Auth), new(*controller.AuthController)),
	provideAuthController,
)

func provideNoteController(nu domain.NoteUsecase, log logrus.Logger) *controller.NoteController {
	return controller.NewNoteController(nu, log)
}

func provideFolderController(fu domain.FolderUsecase, log logrus.Logger) *controller.FolderController {
	return controller.NewFolderController(fu, log)
}

func provideAuthController(uu domain.UserUsecase, log logrus.Logger) *controller.AuthController {
	return controller.NewAuthController(uu, log)
}
