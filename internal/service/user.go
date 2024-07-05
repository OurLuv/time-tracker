package service

import (
	"context"
	"fmt"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/OurLuv/time-tracker/internal/storage"
)

type UserService interface {
	Create(context.Context, model.User) (int, error)
	GetOrderBy(context.Context, string) ([]model.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, u model.User) error
	GetUserByPassport(ctx context.Context, pass string) (*model.User, error)
}

type UserServiceImpl struct {
	repo *storage.Storage
}

func (s *UserServiceImpl) Create(ctx context.Context, user model.User) (int, error) {
	return s.repo.UserStorage.Create(ctx, user)
}

// * Getting users by order
func (s *UserServiceImpl) GetOrderBy(ctx context.Context, param string) ([]model.User, error) {
	fmt.Println(param)
	switch param {
	case "id":
		return s.repo.UserStorage.GetAllUsersOrderById(ctx)
	case "passport_number":
		return s.repo.UserStorage.GetAllUsersOrderByPassportNumber(ctx)
	case "name":
		return s.repo.UserStorage.GetAllUsersOrderByName(ctx)
	case "surname":
		return s.repo.UserStorage.GetAllUsersOrderBySurname(ctx)
	case "patronymic":
		return s.repo.UserStorage.GetAllUsersOrderByPatronymic(ctx)
	case "address":
		return s.repo.UserStorage.GetAllUsersOrderByAddress(ctx)
	default:
		return nil, fmt.Errorf("invalid param")
	}
}

// * Deleting user
func (s *UserServiceImpl) DeleteUser(ctx context.Context, id int) error {
	return s.repo.UserStorage.DeleteUser(ctx, id)
}

// * Updating user
func (s *UserServiceImpl) UpdateUser(ctx context.Context, u model.User) error {
	return s.repo.UserStorage.UpdateUser(ctx, u)
}

func (s *UserServiceImpl) GetUserByPassport(ctx context.Context, pass string) (*model.User, error) {
	return s.repo.UserStorage.GetUserByPassport(ctx, pass)
}

func NewUserServiceImpl(repo *storage.Storage) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}
