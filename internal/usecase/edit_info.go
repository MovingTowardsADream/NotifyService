package usecase

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/repository"
	"context"
)

type EditUseCase struct {
	repo repository.UsersData
}

func NewEditUseCase(repo repository.UsersData) *EditUseCase {
	return &EditUseCase{repo: repo}
}

func (s *EditUseCase) EditPreferences(ctx context.Context, preferences entity.RequestPreferences) error {
	return s.repo.EditUserPreferences(ctx, preferences)
}
