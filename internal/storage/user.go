package storage

import (
	"context"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage interface {
	Create(context.Context, model.User) (int, error)
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func (r *UserRepository) Create(ctx context.Context, user model.User) (int, error) {
	query := "INSERT INTO users (passport_number, name, surname, patronymic, address) VALUES " +
		"($1, $2, $3, $4, $5) RETURNING id"
	var id int
	row := r.pool.QueryRow(ctx, query, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return 1, nil
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
