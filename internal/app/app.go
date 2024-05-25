package app

import (
	"NotifiService/configs"
	v1 "NotifiService/internal/controller/http/v1"
	"NotifiService/internal/gateway"
	"NotifiService/internal/repository"
	"NotifiService/internal/usecase"
	"NotifiService/pkg/httpserver"
	"NotifiService/pkg/postgres"
	"NotifiService/pkg/rabbitmq/rmq_rpc/client"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type App struct {
	HTTPServer *httpserver.Server
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

	gateway := gateway.NewGateway(rmqClient)

	repo := repository.NewRepository(pg)

	useCase := usecase.NewUseCase(repo, gateway)

	// Init http server
	handler := gin.New()
	v1.NewRouter(handler, log, useCase)
	httpServer := httpserver.New(log, handler, httpserver.Port(cfg.HTTP.Port), httpserver.WriteTimeout(cfg.HTTP.Timeout))

	return &App{
		HTTPServer: httpServer,
		DB:         pg,
	}
}
