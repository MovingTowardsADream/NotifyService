package usecase

import (
	"NotifiService/internal/entity"
	"context"
	"fmt"
	"time"
)

type (
	NotifyWorker interface {
		CreateNewNotify(ctx context.Context, notify entity.Notify) error
	}

	NotifyWorkerRepo interface {
		CreateNewNotify(ctx context.Context, user *entity.Notify) error
	}
)

type NotifyWorkerUseCase struct {
}

func NewNotifyWorker() *NotifyWorkerUseCase {
	return &NotifyWorkerUseCase{}
}

func (uc *NotifyWorkerUseCase) CreateNewNotify(ctx context.Context, notify entity.Notify) error {
	time.Sleep(10 * time.Second)
	fmt.Println("SEND MESSAGE ON", notify)

	return nil
}
