package usecase

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/gateway"
	"NotifiService/internal/repository"
	"context"
	"fmt"
	"time"
)

const (
	_defaultTimeout = 5 * time.Second
)

type SendUseCase struct {
	gateway gateway.NotifyGateway
	repo    repository.UsersData
	timeout time.Duration
}

// New -.
func NewSendUseCase(gw gateway.NotifyGateway, repo repository.UsersData) *SendUseCase {
	uc := &SendUseCase{
		gateway: gw,
		repo:    repo,
		timeout: _defaultTimeout,
	}

	return uc
}

func (uc *SendUseCase) SendNotifyForUser(ctx context.Context, notifyRequest entity.RequestNotification) error {
	// Установка timeout на операцию
	ctxTimeout, cancel := context.WithTimeout(ctx, _defaultTimeout)
	defer cancel()

	users_communications, err := uc.repo.GetUserCommunication(ctxTimeout, notifyRequest.UserID)

	if err != nil {
		return fmt.Errorf("SendUseCase - SendNotifyForUsers - uc.repo.GetUserCommunication: %w", err)
	}

	err = uc.gateway.CreateNotifyMessageOnRabbitMQ(ctxTimeout, entity.Notify{
		UserID:     notifyRequest.UserID,
		NotifyType: notifyRequest.NotifyType,
		EmailBody: entity.EmailBody{
			Email:   users_communications.Email,
			Channel: notifyRequest.Channels.EmailChannel,
		},
		PhoneBody: entity.PhoneBody{
			Phone:   users_communications.Phone,
			Channel: notifyRequest.Channels.PhoneChannel,
		},
	})
	if err != nil {
		return fmt.Errorf("SendUseCase - SendNotifyForUsers - uc.gateway.CreateNotifyMessageOnRabbitMQ: %w", err)
	}

	return nil
}
