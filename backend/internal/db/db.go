package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/joho/godotenv"
)

func New() (*bun.DB, error){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, relying on system env")
	}
	dsn := os.Getenv(("DATABASE_URL"))
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	return db, nil
}