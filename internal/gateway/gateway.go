package gateway

import (
	"NotifiService/internal/entity"
	gateway_rabbitmq "NotifiService/internal/gateway/rabbitmq"
	"context"
)

type NotifyGateway interface {
	CreateNotifyMessageOnRabbitMQ(ctx context.Context, communication entity.UserCommunication) error
}

type Gateway struct {
	NotifyGateway
}

func NewGateway(rmq gateway_rabbitmq.NotifyGatewayRMQ) *Gateway {
	return &Gateway{
		NotifyGateway: gateway_rabbitmq.New(rmq),
	}
}
