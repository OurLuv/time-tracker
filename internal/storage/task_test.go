package storage

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/OurLuv/time-tracker/internal/config"
	"github.com/OurLuv/time-tracker/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	cfg := config.Config{
		ServerPort:   "",
		User:         "postgres",
		Password:     "admin",
		DatabaseName: "time_tracker",
		DBPort:       "5432",
	} // fix it
	conn, err = NewPostgresPool(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("failed to init storage: %s", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestFinishTask(t *testing.T) {
	repo := NewTaskStorage(conn)
	_, err := repo.FinishTask(context.Background(), 4)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllTasks(t *testing.T) {
	repo := NewTaskStorage(conn)
	p := model.TaskPeriod{
		//From:   time.Date(2020, 2, 2, 0, 0, 0, 0, time.UTC),
		From:   time.Now(),
		To:     time.Now(),
		UserId: 1,
	}
	_, err := repo.GetTasksByUserIdForPeriod(context.Background(), p)
	if err != nil {
		t.Error(err)
	}
}
