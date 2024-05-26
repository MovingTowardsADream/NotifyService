package amqp_rpc

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/notifyWorkers/usecase"
	rmq_server "NotifiService/pkg/rabbitmq/rmq_rpc/server"
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type notifyWorkerRoutes struct {
	w usecase.NotifyWorker
}

func newNotifyWorkerRoutes(routes map[string]rmq_server.CallHandler, w usecase.NotifyWorker) {
	r := &notifyWorkerRoutes{w}
	{
		routes["createNewNotify"] = r.createNewNotify()
	}
}

func (r *notifyWorkerRoutes) createNewNotify() rmq_server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		var request entity.Notify

		if err := json.Unmarshal(d.Body, &request); err != nil {
			return nil, fmt.Errorf("amqp_rpc - notifyWorkerRoutes - createNewNotify - json.Unmarshal: %w", err)
		}

		err := r.w.CreateNewNotify(context.Background(), request)
		if err != nil {
			return request, fmt.Errorf("amqp_rpc - notifyWorkerRoutes - createNewNotify - r.w.CreateNewNotify: %w", err)
		}

		return nil, nil
	}
}
