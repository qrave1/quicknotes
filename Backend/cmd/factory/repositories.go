package factory

import (
	"context"
	"database/sql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/qrave1/quicknotes/internal/infrastructure/repository"
	"github.com/qrave1/quicknotes/internal/usecase/repositories"
	"github.com/redis/go-redis/v9"
)

var repositorySet = wire.NewSet(
	mustProvideDB,
	mustProvideRedis,

	wire.Bind(new(repository.AuthToken), new(*repository.AuthTokenRedisRepository)),
	provideTokenRepository,

	wire.Bind(new(repositories.Note), new(*repository.NotePostgresRepository)),
	provideNoteRepository,

	wire.Bind(new(repositories.Folder), new(*repository.FolderPostgresRepository)),
	provideFolderRepository,

	wire.Bind(new(repositories.User), new(*repository.UserPostgresRepository)),
	provideUserRepository,
)

func mustProvideDB(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DbConfig.DSN)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func mustProvideRedis(cfg *config.Config) *redis.Client {
	opts, err := redis.ParseURL(cfg.CacheConfig.DSN)
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(opts)

	err = cli.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return cli
}

func provideTokenRepository(c *redis.Client) *repository.AuthTokenRedisRepository {
	return repository.NewTokenRedisRepository(c)
}

func provideNoteRepository(db *sql.DB) *repository.NotePostgresRepository {
	return repository.NewNotePostgresRepository(db)
}

func provideFolderRepository(db *sql.DB) *repository.FolderPostgresRepository {
	return repository.NewFolderPostgresRepository(db)
}

func provideUserRepository(db *sql.DB) *repository.UserPostgresRepository {
	return repository.NewUserPostgresRepository(db)
}
