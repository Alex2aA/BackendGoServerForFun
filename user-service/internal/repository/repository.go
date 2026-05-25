package repository

import (
	"context"

	"github.com/Alex2aA/user-service/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	UpdateRefreshToken(ctx context.Context, userID, token string) error
}
