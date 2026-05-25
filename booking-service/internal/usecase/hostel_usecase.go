package usecase

import (
	"context"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/internal/repository"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type HostelUsecase struct {
	repo repository.HostelRepository
}

func NewHostelUsecase(repo repository.HostelRepository) *HostelUsecase {
	return &HostelUsecase{repo: repo}
}

func (u *HostelUsecase) Create(ctx context.Context, hostel *domain.Hostel) error {
	logger.Log.Info("Creating new hostel", zap.String("name", hostel.Name))

	hostel.ID = uuid.New().String()

	if err := u.repo.Create(ctx, hostel); err != nil {
		logger.Log.Error("Failed to create hostel", zap.Error(err))
		return err
	}

	logger.Log.Info("Hostel created successfully", zap.String("hostel_id", hostel.ID))
	return nil
}