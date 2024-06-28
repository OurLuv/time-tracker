package service

import "github.com/OurLuv/time-tracker/internal/storage"

type Service struct {
	UserService
}

func NewService(repo *storage.Storage) *Service {
	return &Service{
		UserService: NewUserServiceImpl(repo),
	}
}
