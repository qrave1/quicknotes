// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package factory

import (
	"database/sql"
	"github.com/qrave1/quicknotes/internal/infrastructure/interfaces/http"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func InitializeService() (NoteService, func(), error) {
	config := provideConfig()
	validator := provideValidator()
	logger := provideLogger()
	db := mustProvideDB(config, logger)
	userPostgresRepository := provideUserRepository(db)
	client := mustProvideRedis(config, logger)
	authTokenRedisRepository := provideTokenRepository(client)
	authService := provideAuthUsecase(config, logger)
	userService := provideUserUsecase(userPostgresRepository, authTokenRedisRepository, authService, logger)
	authController := provideAuthController(userService, logger)
	folderPostgresRepository := provideFolderRepository(db)
	folderService := provideFolderUsecase(folderPostgresRepository, logger)
	folderController := provideFolderController(folderService, logger)
	notePostgresRepository := provideNoteRepository(db)
	noteService := provideNoteUsecase(notePostgresRepository)
	noteController := provideNoteController(noteService, logger)
	noteServer, cleanup, err := provideServer(config, validator, authController, folderController, noteController)
	if err != nil {
		return NoteService{}, nil, err
	}
	factoryNoteService := provideNoteService(noteServer)
	return factoryNoteService, func() {
		cleanup()
	}, nil
}

func InitializeMigrationContainer() (MigrationContainer, func(), error) {
	config := provideConfig()
	logger := provideLogger()
	db := mustProvideDB(config, logger)
	migrationContainer := provideMigrationContainer(db)
	return migrationContainer, func() {
	}, nil
}

// wire.go:

type NoteService struct{}

func provideNoteService(_ *http.NoteServer) NoteService {
	return NoteService{}
}

type MigrationContainer struct {
	db *sql.DB
}

func (c MigrationContainer) DB() *sql.DB {
	return c.db
}

func provideMigrationContainer(db *sql.DB) MigrationContainer {
	mig := MigrationContainer{db: db}
	return mig
}