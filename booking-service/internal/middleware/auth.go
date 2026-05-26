package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Alex2aA/booking-service/internal/repository"
	"github.com/Alex2aA/booking-service/internal/usecase/token"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	tokenService *token.TokenService
	tokenRepo    repository.TokenRepository
}

func NewAuthMiddleware(ts *token.TokenService, tr repository.TokenRepository) *AuthMiddleware {
	return &AuthMiddleware{tokenService: ts, tokenRepo: tr}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if tokenStr == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := m.tokenService.ParseToken(tokenStr)
		if err != nil {
			logger.Log.Warn("Invalid token", zap.Error(err))
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		blacklisted, _ := m.tokenRepo.IsBlacklisted(r.Context(), tokenStr)
		if blacklisted {
			http.Error(w, "Token blacklisted", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
