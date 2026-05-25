package usecase

import (
	"context"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/internal/repository"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type HouseUsecase struct {
	repo repository.HouseRepository
}

func NewHouseUsecase(repo repository.HouseRepository) *HouseUsecase {
	return &HouseUsecase{repo: repo}
}

func (u *HouseUsecase) Create(ctx context.Context, house *domain.House) error {
	logger.Log.Info("Creating new house", zap.String("name", house.Name))

	house.ID = uuid.New().String()

	if err := u.repo.Create(ctx, house); err != nil {
		logger.Log.Error("Failed to create house", zap.Error(err))
		return err
	}

	logger.Log.Info("House created", zap.String("house_id", house.ID))
	return nil
}