package main

import (
	"fmt"
	"github.com/qrave1/quicknotes/cmd/commands"
	_ "github.com/qrave1/quicknotes/cmd/commands/migrate"
	"github.com/qrave1/quicknotes/cmd/factory"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &cli.App{
		Name: "app",
		Action: func(c *cli.Context) error {
			_, cleanup, err := factory.InitializeService()
			if err != nil {
				log.Fatal(err)
			}
			defer cleanup()

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			<-sigCh

			return nil
		},
		Commands: commands.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		panic(fmt.Errorf("error start app. %w", err))
	}
}
