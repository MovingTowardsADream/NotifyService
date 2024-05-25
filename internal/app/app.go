package app

import (
	"NotifiService/configs"
	"NotifiService/internal/repository"
	"NotifiService/pkg/postgres"
	"fmt"
	"log/slog"
)

type App struct {
	//HTTPServer *httpserver.Server
	//RMQServer  *rmq_server.Server
	DB *postgres.Postgres
}

func New(log *slog.Logger, cfg *configs.Config) *App {

	// Connect postgres db
	pg, err := postgres.NewPostgresDB(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		panic("app - Run - postgres.NewPostgresDB: " + err.Error())
	}

	repo := repository.NewRepository(pg)

	fmt.Println(repo)

	return &App{
		//HTTPServer: ,
		//RMQServer:  ,
		DB: pg,
	}
}
