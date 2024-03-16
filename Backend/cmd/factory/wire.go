//go:build wireinject
// +build wireinject

package factory

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http"
)

type NoteService struct{}

func InitializeService() (NoteService, func(), error) {
	panic(
		wire.Build(
			provideConfig,
			provideLogger,
			provideNoteService,
			controllerSet,
			usecaseSet,
			interfacesSet,
			repositorySet,
		),
	)
}

func provideNoteService(_ *http.NoteServer) NoteService {
	return NoteService{}
}

type MigrationContainer struct {
	db *sql.DB
}

func (c MigrationContainer) DB() *sql.DB {
	return c.db
}

func InitializeMigrationContainer() (MigrationContainer, func(), error) {
	panic(
		wire.Build(
			provideConfig,
			mustProvideDB,
			provideMigrationContainer,
		),
	)
}

func provideMigrationContainer(db *sql.DB) MigrationContainer {
	mig := MigrationContainer{db: db}
	return mig
}
