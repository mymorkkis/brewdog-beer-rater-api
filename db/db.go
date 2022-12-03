package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var databaseURL = fmt.Sprintf(
	"postgres://%s:%s@db:5432/%s",
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_DB"),
)

func connect() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), databaseURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return dbpool
}

func QueryRows[T Model](query string, parser func(pgx.Rows) []T, args ...any) []T {
	dbpool := connect()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), query, args...)
	if err != nil {
		log.Fatalf("Unable to execute ratings query: %v", err)
	}

	return parser(rows)
}
