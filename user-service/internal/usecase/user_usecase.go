package usecase

import (
	"context"

	"github.com/Alex2aA/user-service/internal/domain"
	"github.com/Alex2aA/user-service/internal/repository"
	"github.com/Alex2aA/user-service/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo         repository.UserRepository
	tokenService *TokenService
}

func NewUserUsecase(repo repository.UserRepository, ts *TokenService) *UserUsecase {
	return &UserUsecase{repo: repo, tokenService: ts}
}

func (u *UserUsecase) Register(ctx context.Context, username, password string) (string, error) {
	logger.Log.Info("Registration attempt", zap.String("username", username))

	existing, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if existing != nil {
		logger.Log.Warn("User already exists", zap.String("username", username))
		return "", domain.ErrUserAlreadyExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	user := &domain.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hashed),
		Alias:    "none",
	}

	refreshToken, _ := u.tokenService.GenerateRefreshToken(user.ID)
	user.RefreshToken = refreshToken

	if err := u.repo.Create(ctx, user); err != nil {
		logger.Log.Error("Failed to save user", zap.Error(err))
		return "", err
	}

	accessToken, _ := u.tokenService.GenerateAccessToken(user.ID)
	logger.Log.Info("User registered successfully", zap.String("user_id", user.ID))

	return accessToken, nil
}

func (u *UserUsecase) Login(ctx context.Context, username, password string) (string, error) {
	logger.Log.Info("Login attempt", zap.String("username", username))

	user, err := u.repo.GetByUsername(ctx, username)
	if err != nil || user == nil {
		return "", domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.Log.Warn("Invalid password", zap.String("username", username))
		return "", domain.ErrInvalidCredentials
	}

	accessToken, _ := u.tokenService.GenerateAccessToken(user.ID)
	logger.Log.Info("Login successful", zap.String("user_id", user.ID))

	return accessToken, nil
}
