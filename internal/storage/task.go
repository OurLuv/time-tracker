package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskStorage interface {
	StartTask(ctx context.Context, task model.Task) (int, error)
	FinishTask(ctx context.Context, id int) (*model.Task, error)
	GetTasksByUserIdForPeriod(ctx context.Context, period model.TaskPeriod) ([]model.Task, error)
}

type TaskRepository struct {
	pool *pgxpool.Pool
}

var (
	ErrAlreadyFinished = errors.New("task is already finished")
)

func (r *TaskRepository) StartTask(ctx context.Context, task model.Task) (int, error) {
	query := "INSERT INTO tasks (title, started_at, user_id) VALUES ($1, $2, $3) RETURNING id"
	var id int
	if err := r.pool.QueryRow(ctx, query, task.Title, time.Now(), task.UserId).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskRepository) FinishTask(ctx context.Context, id int) (*model.Task, error) {

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	var isFinished bool
	query := "SELECT is_finished FROM tasks WHERE id = $1"
	if err := tx.QueryRow(ctx, query, id).Scan(&isFinished); err != nil {
		return nil, err
	}
	if isFinished {
		return nil, fmt.Errorf("for task[id][%d]: %w", id, ErrAlreadyFinished)
	}

	query = "UPDATE tasks SET finished_at = $1, is_finished=true WHERE id = $2 RETURNING started_at, finished_at"
	var task model.Task
	if err := tx.QueryRow(ctx, query, time.Now(), id).Scan(&task.StartedAt, &task.FinishedAt); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetTasksByUserIdForPeriod(ctx context.Context, period model.TaskPeriod) ([]model.Task, error) {
	var tasks []model.Task
	var task model.Task

	//2024-05-03
	query := "SELECT * FROM tasks WHERE " +
		"started_at > $1 AND " +
		"finished_at < $2 AND " +
		"user_id = $3 " +
		"ORDER BY finished_at - started_at DESC;"

	rows, err := r.pool.Query(ctx, query, period.From, period.To, period.UserId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Title, &task.StartedAt, &task.FinishedAt, &task.IsFinished, &task.UserId); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func NewTaskStorage(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		pool: pool,
	}
}
