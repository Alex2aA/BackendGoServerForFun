package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alex2aA/booking-service/config"
	deliveryhttp "github.com/Alex2aA/booking-service/internal/delivery/http"
	"github.com/Alex2aA/booking-service/internal/repository/postgres"
	"github.com/Alex2aA/booking-service/internal/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	//_ "github.com/Alex2aA/booking-service/docs"
)

func main() {
	logger.Init()
	defer logger.Sync()

	cfg := config.Load()

	// Миграции
	if err := runMigrations(cfg.DBURL); err != nil {
		logger.Log.Warn("Migration warning", zap.Error(err))
	}

	// Подключение к БД
	db, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		logger.Log.Fatal("Database connection failed", zap.Error(err))
	}
	defer db.Close()

	// Репозитории
	hostelRepo := postgres.NewHostelRepository(db)
	houseRepo := postgres.NewHouseRepository(db)
	bookingRepo := postgres.NewBookingRepository(db)

	// Usecases
	hostelUsecase := usecase.NewHostelUsecase(hostelRepo)
	houseUsecase := usecase.NewHouseUsecase(houseRepo)
	bookingUsecase := usecase.NewBookingUsecase(bookingRepo, houseRepo)

	// Handlers
	hostelHandler := deliveryhttp.NewHostelHandler(hostelUsecase)
	houseHandler := deliveryhttp.NewHouseHandler(houseUsecase)
	bookingHandler := deliveryhttp.NewBookingHandler(bookingUsecase)

	// Router
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api").Subrouter()

	// Public routes (если нужны)
	// Protected routes (в будущем добавить AuthMiddleware)

	api.HandleFunc("/hostel", hostelHandler.Create).Methods("POST")
	api.HandleFunc("/house", houseHandler.Create).Methods("POST")
	api.HandleFunc("/booking", bookingHandler.Create).Methods("POST")

	logger.Log.Info("🚀 Booking-Service started", zap.String("port", cfg.ServerPort))

	srv := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: r,
	}

	// Graceful shutdown
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
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Error("Shutdown error", zap.Error(err))
	}
}

func runMigrations(dbURL string) error {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		return err
	}
	defer m.Close()
	return m.Up()
}