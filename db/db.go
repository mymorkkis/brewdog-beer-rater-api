package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mymorkkis/brewdog-beer-rater-api/internal"
)

type Rows interface {
	Next() bool
	Scan(...any) error
}

func QueryRows[T Model](query string, parser func(Rows) ([]T, error), args ...any) ([]T, error) {
	dbpool, err := connect()
	if err != nil {
		return nil, err
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), query, args...)
	if err != nil {
		internal.ErrorLog.Printf("Unable to execute ratings query: %v", err)
		return nil, err
	}

	return parser(rows)
}

func connect() (*pgxpool.Pool, error) {
	connectionURL := fmt.Sprintf(
		"postgres://%s:%s@db:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	dbpool, err := pgxpool.New(context.Background(), connectionURL)

	if err != nil {
		internal.ErrorLog.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	return dbpool, nil
}
