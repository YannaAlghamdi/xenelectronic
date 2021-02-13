package main

import (
	"log"
	"os"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := config.NewConfig("./config/config.yaml")

	if err != nil {
		panic(err)
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "Migration id to migrate up to or to rollback down to.",
		},
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Run server.",
				Action: func(c *cli.Context) error {
					server, err := buildServer(cfg)
					if err != nil {
						panic(err)
					}
					server.Start()
					return nil
				},
			},
			{
				Name:  "migrate",
				Usage: "Run all migrations or migrate to specfic migration ID.",
				Action: func(c *cli.Context) error {
					migration, err := initializeMigration(cfg)
					if err != nil {
						panic(err)
					}

					if c.String("id") != "" {
						err = migration.Migrate(c.String("id"))
					} else {
						err = migration.Migrate("")
					}

					return err
				},
				Flags: flags,
			},
			{
				Name:  "rollback",
				Usage: "Run all rollbacks or a rollback to specific migration ID",
				Action: func(c *cli.Context) error {
					migration, err := initializeMigration(cfg)
					if err != nil {
						panic(err)
					}
					err = migration.Rollback(c.String("id"))
					return err
				},
				Flags: flags,
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
