package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/chagasVinicius/apollo/cmd/tooling/admin/commands"
	"github.com/chagasVinicius/apollo/internal/sys/database"
	"go.uber.org/zap"
)

var build = "develop"

type config struct {
	conf.Version

	Args conf.Args
	DB   struct {
		User       string `conf:"default:postgres"`
		Password   string `conf:"default:postgres,mask"`
		Host       string `conf:"default:localhost"`
		Name       string `conf:"default:postgres"`
		DisableTLS bool   `conf:"default:true"`
	}
}

func main() {
	if err := run(zap.NewNop().Sugar()); err != nil {
		if !errors.Is(err, commands.ErrHelp) {
			fmt.Println("ERROR", err)
		}
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	cfg := config{
		Version: conf.Version{
			Build: build,
		},
	}

	const prefix = "APOLLO"

	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		out, err := conf.String(&cfg)
		if err != nil {
			return fmt.Errorf("generating config for output: %w", err)
		}
		log.Infow("startup", "config", out)

		return fmt.Errorf("parsing config: %w", err)
	}

	return processCommands(cfg)
}

// processCommands handles the execution of the commands specified on
// the command line.
func processCommands(cfg config) error {
	dbConfig := database.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	}
	fmt.Println(dbConfig)
	switch cfg.Args.Num(0) {
	case "migrate":
		if err := commands.Migrate(dbConfig); err != nil {
			return fmt.Errorf("migrating database: %w", err)
		}

	default:
		fmt.Println("migrate:    create the schema in the database")
		fmt.Println("provide a command to get more help.")
		return commands.ErrHelp
	}
	return nil
}
