package main

import (
	"context"
	"log"

	"github.com/promingy/yelp-clone/backend/internal/db"
	"github.com/promingy/yelp-clone/backend/internal/migrations"
	"github.com/uptrace/bun/migrate"
)

func main() {
    ctx := context.Background()
    database, err := db.New()
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    migrator := migrate.NewMigrator(database, migrations.Migrations)

    if err := migrator.Init(ctx); err != nil {
        log.Fatal(err)
    }

    if _, err := migrator.Migrate(ctx); err != nil {
        log.Fatal(err)
    }

    log.Println("Migrations applied successfully")
}