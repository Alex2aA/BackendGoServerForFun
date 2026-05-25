package postgres

import (
	"context"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type BookingRepository struct {
	db *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Create(ctx context.Context, booking *domain.Booking) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO booked (id, id_house, id_hostel, date_start, date_end)
		VALUES ($1, $2, $3, $4, $5)`,
		booking.ID, booking.HouseID, booking.HostelID, booking.DateStart, booking.DateEnd)

	if err != nil {
		logger.Log.Error("Failed to create booking", zap.Error(err))
	}
	return err
}