package service

import (
	"context"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/OurLuv/time-tracker/internal/storage"
)

type UserService interface {
	Create(context.Context, model.User) (int, error)
}

type UserServiceImpl struct {
	repo *storage.Storage
}

func (s *UserServiceImpl) Create(ctx context.Context, user model.User) (int, error) {
	return s.repo.UserStorage.Create(ctx, user)
}

func NewUserServiceImpl(repo *storage.Storage) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}
