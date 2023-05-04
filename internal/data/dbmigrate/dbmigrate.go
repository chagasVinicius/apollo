// Package dbmigrate contains the database schema, migrations and seeding data.
package dbmigrate

import (
	"context"
	"embed"
	"fmt"

	"github.com/chagasVinicius/apollo/internal/sys/database"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// Migrations is a list of database Migrations.
var Migrations = migrate.NewMigrations()

//go:embed sql
var sqlMigrations embed.FS

func init() {
	// Must discovery all migrations.
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(fmt.Errorf("failed to discover database migrations: %w", err))
	}
}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, db *bun.DB) error {
	if err := database.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("database is not read: %w", err)
	}

	m := migrate.NewMigrator(db, Migrations)

	if err := m.Init(ctx); err != nil {
		return fmt.Errorf("failed to initialize migrator tables: %w", err)
	}

	if _, err := m.Migrate(ctx); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
