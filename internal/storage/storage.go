package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	UserStorage
	TaskStorage
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{
		UserStorage: NewUserRepository(pool),
		TaskStorage: NewTaskStorage(pool),
	}
}
