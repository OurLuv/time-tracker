package main

import (
	"log/slog"
	"os"

	"github.com/OurLuv/time-tracker/internal/config"
	"github.com/OurLuv/time-tracker/internal/handler"
	"github.com/phsym/console-slog"
)

func main() {

	// getting config data
	cfg := config.MustLoad()

	// logger
	log := SetupLogger()
	log.Debug("Starting app with config", slog.Any("cfg", cfg))

	// starting server
	log.Info("Starting server", slog.String("port", cfg.ServerPort))
	h := handler.Handler{}
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
