package usecase

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/gateway"
	"NotifiService/internal/repository"
	"context"
)

type EditInfo interface {
	EditPreferences(ctx context.Context, preferences entity.UserPreferences) error
}

type NotifySend interface {
	SendNotifyForUser(ctx context.Context, id string) error
}

type UseCase struct {
	EditInfo
	NotifySend
}

func NewUseCase(repos *repository.Repository, gw gateway.Gateway) *UseCase {
	return &UseCase{
		EditInfo:   NewEditUseCase(repos.UsersData),
		NotifySend: NewSendUseCase(gw.NotifyGateway, repos.UsersData),
	}
}
