package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connect() *pgxpool.Pool {
	database_url := fmt.Sprintf(
		"postgres://%s:%s@db:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	dbpool, err := pgxpool.New(context.Background(), database_url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}

func greetingTest(dbpool *pgxpool.Pool) {
	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'DB query working'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
