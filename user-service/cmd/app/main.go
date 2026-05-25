package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alex2aA/user-service/config"
	_ "github.com/Alex2aA/user-service/docs"
	deliveryhttp "github.com/Alex2aA/user-service/internal/delivery/http"
	"github.com/Alex2aA/user-service/internal/repository/postgres"
	"github.com/Alex2aA/user-service/internal/usecase"
	"github.com/Alex2aA/user-service/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.Sync()

	cfg := config.Load()

	if err := runMigrations(cfg.DBURL); err != nil {
		logger.Log.Warn("Migration warning", zap.Error(err))
	}

	db, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		logger.Log.Fatal("DB connection failed", zap.Error(err))
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	tokenService := usecase.NewTokenService(cfg.JWTSecret)
	userUsecase := usecase.NewUserUsecase(userRepo, tokenService)

	userHandler := deliveryhttp.NewUserHandler(userUsecase)

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/user/login", userHandler.Login).Methods("POST")

	logger.Log.Info("🚀 User-Service started", zap.String("port", cfg.ServerPort))

	srv := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Server crashed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func runMigrations(dbURL string) error {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		return err
	}
	defer m.Close()
	return m.Up()
}
