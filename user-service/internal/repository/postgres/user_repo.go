package postgres

import (
	"context"
	"database/sql"

	"github.com/Alex2aA/user-service/internal/domain"
	"github.com/Alex2aA/user-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users (id, username, hash_password, alias, refresh_token)
		VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.Username, user.Password, user.Alias, user.RefreshToken)

	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err), zap.String("username", user.Username))
	}
	return err
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(ctx, `
		SELECT id, username, hash_password, alias, refresh_token 
		FROM users WHERE username = $1`, username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Alias, &user.RefreshToken)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logger.Log.Error("Failed to get user", zap.Error(err))
	}
	return user, err
}

func (r *UserRepository) UpdateRefreshToken(ctx context.Context, userID, token string) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET refresh_token = $1 WHERE id = $2", token, userID)
	return err
}
