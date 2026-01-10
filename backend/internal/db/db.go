package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func New() (*bun.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("No .env file found, relying on system env.")
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case "dev":
		dsn := os.Getenv("SQLITE_DSN")
		if dsn == "" {
			return nil, fmt.Errorf("SQLITE_DSN env is not set.")
		}

		sqldb, err := sql.Open(sqliteshim.ShimName, dsn)
		if err != nil {
			return nil, err
		}
		db := bun.NewDB(sqldb, sqlitedialect.New())

		return db, nil

	case "prod":
		dsn := os.Getenv("POSTGRES_DSN")
		if dsn == "" {
			return nil, fmt.Errorf("POSTGRES_DSN env is not set.")
		}
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		db := bun.NewDB(sqldb, pgdialect.New())
		
		return db, nil
	default:
		return nil, fmt.Errorf("Unknown APP_ENV: %s", env)
	}
}
