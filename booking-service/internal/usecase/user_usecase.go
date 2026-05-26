package usecase

import (
	"context"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/internal/repository"
	"github.com/Alex2aA/booking-service/internal/usecase/token"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserUsecase struct {
	repo         repository.UserRepository
	tokenService *token.TokenService
}

func NewUserUsecase(repo repository.UserRepository, ts *token.TokenService) *UserUsecase {
	return &UserUsecase{repo: repo, tokenService: ts}
}

func (u *UserUsecase) Register(ctx context.Context, username, password string) (string, error) {
	logger.Log.Info("Registration attempt", zap.String("username", username))

	user := &domain.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: password,
		Alias:    "none",
	}

	if err := u.repo.Create(ctx, user); err != nil {
		logger.Log.Error("Failed to save user", zap.Error(err))
		return "", err
	}

	accessToken, _ := u.tokenService.GenerateAccessToken(user.ID)
	logger.Log.Info("User registered successfully", zap.String("user_id", user.ID))

	return accessToken, nil
}
