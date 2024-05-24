package main

import (
	"NotifiService/configs"
	"NotifiService/internal/app"
	"NotifiService/pkg/logger"
	"fmt"
)

func main() {
	// Init configuration
	cfg := configs.MustLoad()

	// Init logger
	log := logger.SetupLogger(cfg.Log.Level)

	application := app.New(log, cfg)

	fmt.Println(application)
}
