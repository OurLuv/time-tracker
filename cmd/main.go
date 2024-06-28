package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/OurLuv/time-tracker/internal/config"
	"github.com/OurLuv/time-tracker/internal/handler"
	"github.com/OurLuv/time-tracker/internal/service"
	"github.com/OurLuv/time-tracker/internal/storage"
	"github.com/phsym/console-slog"
)

func main() {

	// getting config data
	cfg := config.MustLoad()

	// logger
	log := SetupLogger()
	log.Debug("Starting app with config", slog.Any("cfg", cfg))

	connString := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", cfg.User, cfg.Password, cfg.DBPort, cfg.DatabaseName)
	log.Info(connString)

	// conn to db
	pool, err := storage.NewPostgresPool(context.Background(), cfg)
	defer pool.Close()
	if err != nil {
		log.Error("Panic", slog.String("err", err.Error()))
		os.Exit(1)
	}
	log.Info("Connected to db", slog.String("port", cfg.DBPort))

	// init layers
	repo := storage.NewStorage(pool)
	service := service.NewService(repo)
	h := handler.NewHandler(service, log)

	// starting server
	log.Info("Starting server", slog.String("port", cfg.ServerPort))
	r := h.InitRoutes()
	server := handler.Server{}
	server.Start(cfg.ServerPort, r)
}

func SetupLogger() *slog.Logger {
	logger := slog.New(
		console.NewHandler(os.Stderr, &console.HandlerOptions{Level: slog.LevelDebug}),
	)
	slog.SetDefault(logger)

	return logger
}
