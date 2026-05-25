package postgres

import (
	"context"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HostelRepository struct {
	db *pgxpool.Pool
}

func NewHostelRepository(db *pgxpool.Pool) *HostelRepository {
	return &HostelRepository{db: db}
}

func (r *HostelRepository) Create(ctx context.Context, hostel *domain.Hostel) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO hostels (id, name, description, rate)
		VALUES ($1, $2, $3, $4)`,
		hostel.ID, hostel.Name, hostel.Description, hostel.Rate)

	if err != nil {
		logger.Log.Error("Failed to create hostel", zap.Error(err))
	}
	return err
}