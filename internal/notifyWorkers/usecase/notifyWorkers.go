package usecase

import (
	"NotifiService/internal/entity"
	"context"
	"fmt"
	"time"
)

type NotifyWorker interface {
	CreateNewNotify(ctx context.Context, notify entity.Notify) error
}

type NotifyWorkerUseCase struct {
}

func NewNotifyWorker() *NotifyWorkerUseCase {
	return &NotifyWorkerUseCase{}
}

func (uc *NotifyWorkerUseCase) CreateNewNotify(ctx context.Context, notify entity.Notify) error {
	time.Sleep(1 * time.Second)
	fmt.Println("SEND MESSAGE ON", notify)

	return nil
}
