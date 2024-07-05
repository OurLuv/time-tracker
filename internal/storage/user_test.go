package storage

import (
	"context"
	"testing"
)

func TestGetAllUsersOrderById(t *testing.T) {
	repo := NewUserRepository(conn)
	users, err := repo.GetAllUsersOrderById(context.Background())
	if err != nil {
		t.Error(err)
	}
	_ = users
}

func TestGetAllUsersOrderByName(t *testing.T) {
	repo := NewUserRepository(conn)
	users, err := repo.GetAllUsersOrderByName(context.Background())
	if err != nil {
		t.Error(err)
	}
	_ = users
}
