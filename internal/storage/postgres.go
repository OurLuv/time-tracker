package storage

import (
	"context"
	"fmt"

	"github.com/OurLuv/time-tracker/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", cfg.User, cfg.Password, cfg.DBPort, cfg.DatabaseName)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	var greeting string
	err = pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return nil, err
	}
	fmt.Print(greeting)

	return pool, nil
}
