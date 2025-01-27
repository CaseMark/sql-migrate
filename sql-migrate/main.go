package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "Configuration file to use",
		},
		&cli.StringFlag{
			Name:  "env",
			Usage: "Environment to use",
		},
	}
	app := &cli.App{
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
		Version:   GetVersion(),
		Commands: []*cli.Command{
			{
				Name:        "up",
				Usage:       (&UpCommand{}).Help(),
				Description: (&UpCommand{}).Synopsis(),
				Action: func(cCtx *cli.Context) error {
					(&UpCommand{}).Run(os.Args[2:])
					return nil
				},
				Flags: flags,
			},
			{
				Name:        "down",
				Usage:       (&DownCommand{}).Help(),
				Description: (&DownCommand{}).Synopsis(),
				Action: func(cCtx *cli.Context) error {
					(&DownCommand{}).Run(os.Args[2:])
					return nil
				},
				Flags: flags,
			},
			{
				Name:        "redo",
				Usage:       (&RedoCommand{}).Help(),
				Description: (&RedoCommand{}).Synopsis(),
				Action: func(cCtx *cli.Context) error {
					(&RedoCommand{}).Run(os.Args[2:])
					return nil
				},
				Flags: flags,
			},
			{
				Name:        "status",
				Usage:       (&StatusCommand{}).Help(),
				Description: (&StatusCommand{}).Synopsis(),
				Action: func(cCtx *cli.Context) error {
					(&StatusCommand{}).Run(os.Args[2:])
					return nil
				},
				Flags: flags,
			},
			{
				Name:        "new",
				Usage:       (&NewCommand{}).Help(),
				Description: (&NewCommand{}).Synopsis(),
				Action: func(cCtx *cli.Context) error {
					(&NewCommand{}).Run(os.Args[2:])
					return nil
				},
				Flags: flags,
			},
			{
				Name:        "skip",
				Usage:       (&SkipCommand{}).Help(),
				Description: (&SkipCommand{}).Synopsis(),
				Action: func(cCtx *cli.Context) error {
					(&SkipCommand{}).Run(os.Args[2:])
					return nil
				},
				Flags: flags,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		return 1
	}

	return 0
}
