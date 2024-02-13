package migrate

import (
	"context"
	"github.com/pressly/goose/v3"
	"github.com/qrave1/quicknotes/cmd/commands"
	"github.com/qrave1/quicknotes/internal/storage"
	"github.com/urfave/cli/v2"
	"time"
)

func init() {
	commands.RegisterCommand(&cli.Command{
		Name: "migrate",
		Action: func(c *cli.Context) error {
			// TODO migrate di container

			goose.SetBaseFS(storage.EmbedMigrations)
			err := goose.SetDialect("postgres")
			if err != nil {
				panic(err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			return goose.RunContext(
				ctx,
				c.Args().First(),
				// todo db from di cont here
				db,
				"migrations",
				c.Args().Get(1),
			)
		},
	})
}
