package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/chagasVinicius/apollo/internal/data/dbmigrate"
	"github.com/chagasVinicius/apollo/internal/sys/database"
	"github.com/urfave/cli/v2"
)

func DBCommand(cfg database.Config) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "manage database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					db, err := database.Open(cfg)
					if err != nil {
						return fmt.Errorf("connect database: %w", err)
					}
					defer db.Close()

					return nil
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					db, err := database.Open(cfg)
					if err != nil {
						return fmt.Errorf("connect database: %w", err)
					}
					defer db.Close()

					ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()

					if err := dbmigrate.Migrate(ctx, db); err != nil {
						return fmt.Errorf("migrate database: %w", err)
					}

					fmt.Println("migrations complete")
					return nil
				},
			},
		},
	}
}
