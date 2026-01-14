package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/promingy/yelp-clone/backend/internal/db"
	"github.com/promingy/yelp-clone/backend/internal/migrations"
	"github.com/uptrace/bun/migrate"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Set default environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
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
