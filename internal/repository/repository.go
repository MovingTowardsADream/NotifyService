package repository

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/repository/postgresdb"
	"NotifiService/pkg/postgres"
	"context"
)

type UsersData interface {
	GetUserCommunication(ctx context.Context, id string) (entity.UserCommunication, error)
	EditUserPreferences(ctx context.Context, preferences entity.UserPreferences) error
}

type Repository struct {
	UsersData
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{
		UsersData: postgresdb.NewNotifyRepo(pg),
	}
}
