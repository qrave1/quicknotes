package factory

import (
	"github.com/google/wire"
	"github.com/qrave1/logwrap"
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

func provideNoteController(nu domain.NoteUsecase, log logwrap.Logger) *controller.NoteController {
	return controller.NewNoteController(nu, log)
}

func provideFolderController(fu domain.FolderUsecase, log logwrap.Logger) *controller.FolderController {
	return controller.NewFolderController(fu, log)
}

func provideAuthController(uu domain.UserUsecase, log logwrap.Logger) *controller.AuthController {
	return controller.NewAuthController(uu, log)
}
