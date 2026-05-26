package repository

import (
	"context"

	"github.com/Alex2aA/booking-service/internal/domain"
)

type HostelRepository interface {
	Create(ctx context.Context, hostel *domain.Hostel) error
}

type HouseRepository interface {
	Create(ctx context.Context, house *domain.House) error
	GetByID(ctx context.Context, id string) (*domain.House, error)
}

type BookingRepository interface {
	Create(ctx context.Context, booking *domain.Booking) error
}

type TokenRepository interface {
	IsBlacklisted(ctx context.Context, token string) (bool, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
}
