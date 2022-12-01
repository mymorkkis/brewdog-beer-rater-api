package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var databaseURL = fmt.Sprintf(
	"postgres://%s:%s@db:5432/%s",
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_DB"),
)

func Connect() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}

func GreetingTest(dbpool *pgxpool.Pool) {
	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'DB query working'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
