package app

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type App struct {
  Cli *cli.App
} 

func NewApp() *App {
  app := &App{}

	app.Cli = &cli.App{
		Name:  "greet",
		Usage: "test usage",
		Action: func(*cli.Context) error {
			fmt.Println("hello there")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Action: app.addCmd,
      },
		},
	}

  return app
}

func (a *App) Run() {
  if err := a.Cli.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
