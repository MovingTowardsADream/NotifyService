package app

import (
	"NotifiService/configs"
	"NotifiService/pkg/httpserver"
	"NotifiService/pkg/postgres"
	"log/slog"
)

type App struct {
	HTTPServer *httpserver.Server
	RMQServer  *rmq_server.Server
	DB         *postgres.Postgres
}

func New(log *slog.Logger, cfg *configs.Config) *App {
	// Connect postgres db
	pg, err := postgres.NewPostgresDB(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		panic("app - Run - postgres.NewPostgresDB: " + err.Error())
	}

	return &App{
		HTTPServer: nil,
		RMQServer:  nil,
		DB:         pg,
	}
}
