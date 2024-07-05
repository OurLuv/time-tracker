package service

import "github.com/OurLuv/time-tracker/internal/storage"

type Service struct {
	UserService
	TaskService
}

func NewService(repo *storage.Storage) *Service {
	return &Service{
		UserService: NewUserServiceImpl(repo),
		TaskService: NewTaskServiceImpl(repo),
	}
}
