package service

import (
	"context"
	"fmt"
	"time"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/OurLuv/time-tracker/internal/storage"
)

type TaskService interface {
	StartTask(ctx context.Context, task model.Task) (int, error)
	FinishTask(ctx context.Context, id int) (time.Duration, error)
	GetTasks(ctx context.Context, period model.TaskPeriod) ([]model.Task, error)
}

type TaskServiceImpl struct {
	repo *storage.Storage
}

func (s *TaskServiceImpl) StartTask(ctx context.Context, task model.Task) (int, error) {
	return s.repo.TaskStorage.StartTask(ctx, task)
}

func (s *TaskServiceImpl) FinishTask(ctx context.Context, id int) (time.Duration, error) {
	task, err := s.repo.TaskStorage.FinishTask(ctx, id)
	if err != nil {
		return 0, nil
	}
	dur := task.GetDuration()

	return dur, nil
}

// * Gettimg all tasks of user for certain period
func (s *TaskServiceImpl) GetTasks(ctx context.Context, period model.TaskPeriod) ([]model.Task, error) {
	tasks, err := s.repo.TaskStorage.GetTasksByUserIdForPeriod(ctx, period)
	if err != nil {
		return nil, err
	}
	fmt.Println(tasks)
	for i := range tasks {
		tasks[i].DurationStr = tasks[i].GetDuration().String()
	}

	return tasks, nil
}

func NewTaskServiceImpl(repo *storage.Storage) *TaskServiceImpl {
	return &TaskServiceImpl{
		repo: repo,
	}
}
