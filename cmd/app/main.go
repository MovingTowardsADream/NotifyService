package main

import (
	"NotifiService/configs"
	"NotifiService/pkg/logger"
)

func main() {
	// Init configuration
	cfg := configs.MustLoad()

	// Init logger
	log := logger.SetupLogger(cfg.Log.Level)

	application := app.New(log, cfg)
}
