package factory

import (
	"github.com/google/wire"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http"
	"github.com/qrave1/quicknotes/internal/interface/controller"
	"github.com/qrave1/quicknotes/pkg/validator"
)

var interfacesSet = wire.NewSet(
	provideValidator,
	provideServer,
)

func provideValidator() validator.Validator {
	return validator.New()
}

func provideServer(
	cfg *config.Config,
	validate validator.Validator,
	authController controller.Auth,
	folderController controller.Folder,
	noteController controller.Note,
) (*http.NoteServer, func(), error) {
	s := http.NewNoteServer(cfg, validate, authController, folderController, noteController)

	err := s.Run(cfg)
	if err != nil {
		return nil, nil, err
	}

	return s, s.Shutdown, nil
}
