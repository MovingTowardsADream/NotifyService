package amqp_rpc

import (
	"NotifiService/internal/notifyWorkers/usecase"
	rmq_server "NotifiService/pkg/rabbitmq/rmq_rpc/server"
)

func NewRouter(r usecase.NotifyWorker) map[string]rmq_server.CallHandler {
	routes := make(map[string]rmq_server.CallHandler)
	{
		newNotifyWorkerRoutes(routes, r)
	}

	return routes
}
