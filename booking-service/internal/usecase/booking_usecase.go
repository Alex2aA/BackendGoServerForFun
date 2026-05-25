package usecase

import (
	"context"
	"time"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/internal/repository"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BookingUsecase struct {
	bookingRepo repository.BookingRepository
	houseRepo   repository.HouseRepository
}

func NewBookingUsecase(bookingRepo repository.BookingRepository, houseRepo repository.HouseRepository) *BookingUsecase {
	return &BookingUsecase{
		bookingRepo: bookingRepo,
		houseRepo:   houseRepo,
	}
}

func (u *BookingUsecase) Create(ctx context.Context, houseID, userID, dateStartStr, dateEndStr string) error {
	logger.Log.Info("New booking attempt",
		zap.String("house_id", houseID),
		zap.String("user_id", userID))

	// Проверяем существование дома
	house, err := u.houseRepo.GetByID(ctx, houseID)
	if err != nil || house == nil {
		return domain.ErrHouseNotFound
	}

	// Парсим даты
	dateStart, err := time.Parse("2006-01-02", dateStartStr)
	if err != nil {
		return domain.ErrInvalidDateFormat
	}
	dateEnd, err := time.Parse("2006-01-02", dateEndStr)
	if err != nil {
		return domain.ErrInvalidDateFormat
	}

	if dateEnd.Before(dateStart) {
		return domain.ErrInvalidBookingDates
	}

	// Создаём бронирование
	booking := &domain.Booking{
		ID:        uuid.New().String(),
		HouseID:   houseID,
		HostelID:  house.HostelID,
		DateStart: dateStart,
		DateEnd:   dateEnd,
	}

	if err := u.bookingRepo.Create(ctx, booking); err != nil {
		logger.Log.Error("Failed to create booking", zap.Error(err))
		return err
	}

	logger.Log.Info("Booking created successfully", zap.String("booking_id", booking.ID))
	return nil
}