package app

import (
	"NotifiService/configs"
	"NotifiService/internal/entity"
	"NotifiService/internal/repository"
	"NotifiService/pkg/postgres"
	"context"
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

	pref := entity.UserPreferences{
		UserID: "5dc6900bf5dc29d4685b6bfd63c9d257",
		Preferences: []entity.Preference{
			entity.Preference{
				NotifyType: "promotions",
				Channel:    "email",
				Approval:   true,
			},
		},
	}
	err = repo.UsersData.EditUserPreferences(context.Background(), pref)

	if err != nil {
		fmt.Println(err)
	}

	return &App{
		//HTTPServer: ,
		//RMQServer:  ,
		DB: pg,
	}
}
