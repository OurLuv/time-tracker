package storage

import (
	"context"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage interface {
	Create(context.Context, model.User) (int, error)
	GetAllUsersOrderById(ctx context.Context) ([]model.User, error)
	GetAllUsersOrderByPassportNumber(ctx context.Context) ([]model.User, error)
	GetAllUsersOrderByName(ctx context.Context) ([]model.User, error)
	GetAllUsersOrderBySurname(ctx context.Context) ([]model.User, error)
	GetAllUsersOrderByPatronymic(ctx context.Context) ([]model.User, error)
	GetAllUsersOrderByAddress(ctx context.Context) ([]model.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, u model.User) error
	GetUserByPassport(ctx context.Context, pass string) (*model.User, error)
}

type UserRepository struct {
	pool *pgxpool.Pool
}

// * Create user
func (r *UserRepository) Create(ctx context.Context, user model.User) (int, error) {
	query := "INSERT INTO users (passport_number, name, surname, patronymic, address) VALUES " +
		"($1, $2, $3, $4, $5) RETURNING id"
	var id int
	row := r.pool.QueryRow(ctx, query, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

// * Get all users order by id
func (r *UserRepository) GetAllUsersOrderById(ctx context.Context) ([]model.User, error) {
	var users []model.User
	var user model.User

	query := "SELECT * FROM users"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// * Get all users order by passport number
func (r *UserRepository) GetAllUsersOrderByPassportNumber(ctx context.Context) ([]model.User, error) {
	var users []model.User
	var user model.User

	query := "SELECT * FROM users ORDER BY passport_number"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// * Get all users order by name
func (r *UserRepository) GetAllUsersOrderByName(ctx context.Context) ([]model.User, error) {
	var users []model.User
	var user model.User

	query := "SELECT * FROM users ORDER BY name"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// * Get all users order by surname
func (r *UserRepository) GetAllUsersOrderBySurname(ctx context.Context) ([]model.User, error) {
	var users []model.User
	var user model.User

	query := "SELECT * FROM users ORDER BY surname"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// * Get all users order by patronymic
func (r *UserRepository) GetAllUsersOrderByPatronymic(ctx context.Context) ([]model.User, error) {
	var users []model.User
	var user model.User

	query := "SELECT * FROM users ORDER BY patronymic"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// * Get all users order by address
func (r *UserRepository) GetAllUsersOrderByAddress(ctx context.Context) ([]model.User, error) {
	var users []model.User
	var user model.User

	query := "SELECT * FROM users ORDER BY address"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// * Deleting user
func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := "DELETE FROM users WHERE id = $1"
	if _, err := tx.Exec(ctx, query, id); err != nil {
		return err
	}

	query = "DELETE FROM tasks WHERE user_id= $1"
	if _, err := tx.Exec(ctx, query, id); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// * Change user
func (r *UserRepository) UpdateUser(ctx context.Context, u model.User) error {
	query := "UPDATE users " +
		"SET  " +
		"  passport_number = $1, " +
		"  name = $2, " +
		"  surname = $3, " +
		"  patronymic = $4, " +
		"  address = $5 " +
		"WHERE id = $6;"
	_, err := r.pool.Exec(ctx, query, u.PassportNumber, u.Name, u.Surname, u.Patronymic, u.Address, u.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByPassport(ctx context.Context, pass string) (*model.User, error) {
	query := "SELECT * FROM users WHERE passport_number = $1"
	var user model.User
	if err := r.pool.QueryRow(ctx, query, pass).Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
