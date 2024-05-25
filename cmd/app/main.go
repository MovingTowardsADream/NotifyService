package main

import (
	"NotifiService/configs"
	"NotifiService/internal/app"
	"NotifiService/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Init configuration
	cfg := configs.MustLoad()

	// Init logger
	log := logger.SetupLogger(cfg.Log.Level)

	application := app.New(log, cfg)

	// Run servers
	go func() {
		application.HTTPServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
	}

	log.Info("Starting graceful shutdown")

	if err := application.HTTPServer.Shutdown(); err != nil {
		log.Error("HTTPServer.Shutdown error", logger.Err(err))
	}

	application.DB.Close()

	log.Info("Gracefully stopped")
}
