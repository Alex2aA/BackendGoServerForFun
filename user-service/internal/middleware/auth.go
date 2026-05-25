package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Alex2aA/user-service/internal/usecase"
	"github.com/Alex2aA/user-service/pkg/logger"
	"go.uber.org/zap"
)

func AuthMiddleware(tokenService *usecase.TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			if tokenStr == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := tokenService.ParseToken(tokenStr) // Нужно добавить метод ParseToken в TokenService
			if err != nil {
				logger.Log.Warn("Invalid token", zap.Error(err))
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
