package rabbitmq

import (
	"NotifiService/internal/entity"
	"context"
	"fmt"
)

type NotifyGatewayRMQ interface {
	RemoteCall(ctx context.Context, handler string, request interface{}) error
}

type NotifyGateway struct {
	rmq NotifyGatewayRMQ
}

// Init of wallet gateway, through we will making requests to rmq server.
func New(rmq NotifyGatewayRMQ) *NotifyGateway {
	return &NotifyGateway{rmq}
}

// Creating new wallet with balance, through remote call to rmq server.
func (gw *NotifyGateway) CreateNotifyMessageOnRabbitMQ(ctx context.Context, communication entity.UserCommunication) error {

	err := wrapper(ctx, func() error {
		return gw.rmq.RemoteCall(ctx, "createNewWallet", communication)
	})

	if err != nil {
		return fmt.Errorf("WalletGateway - CreateNewWalletWithBalance - gw.rmq.RemoteCall: %w", err)
	}

	return nil
}

func wrapper(ctx context.Context, f func() error) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- f()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err() //nolint:wrapcheck // we need just a send ctx error
	case err := <-errCh:
		return err
	}
}
