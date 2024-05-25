package app

import (
	"NotifiService/configs"
	"NotifiService/internal/repository"
	"NotifiService/pkg/postgres"
	"NotifiService/pkg/rabbitmq/rmq_rpc/client"
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

	rmqClient, err := client.NewRabbitMQClient(cfg.RMQ.URL, cfg.RMQ.ServerExchange, cfg.RMQ.ClientExchange)
	if err != nil {
		panic("app - Run - rmqServer - server.New" + err.Error())
	}

	repo := repository.NewRepository(pg)

	return &App{
		//HTTPServer: ,
		//RMQServer:  ,
		DB: pg,
	}
}
