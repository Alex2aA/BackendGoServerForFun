package postgres

import (
	"context"
	"database/sql"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"go.uber.org/zap"
)

type HouseRepository struct {
	db *pgxpool.Pool
}

func NewHouseRepository(db *pgxpool.Pool) *HouseRepository {
	return &HouseRepository{db: db}
}

func (r *HouseRepository) Create(ctx context.Context, house *domain.House) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO free_houses (id, id_user, id_hostel, address, number_of_rooms, price_per_day, name, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		house.ID, house.UserID, house.HostelID, house.Address, house.NumberOfRooms,
		house.PricePerDay, house.Name, house.Description)

	if err != nil {
		logger.Log.Error("Failed to create house", zap.Error(err))
	}
	return err
}

func (r *HouseRepository) GetByID(ctx context.Context, id string) (*domain.House, error) {
	h := &domain.House{}
	err := r.db.QueryRow(ctx, `
		SELECT id, id_user, id_hostel, address, number_of_rooms, price_per_day, name, description 
		FROM free_houses WHERE id = $1`, id).
		Scan(&h.ID, &h.UserID, &h.HostelID, &h.Address, &h.NumberOfRooms, &h.PricePerDay, &h.Name, &h.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return h, err
}