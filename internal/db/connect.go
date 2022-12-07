package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenPGXPool() (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@db:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
