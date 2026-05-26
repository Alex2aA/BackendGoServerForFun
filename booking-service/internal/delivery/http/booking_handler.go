package http

import (
	"encoding/json"
	"net/http"

	"github.com/Alex2aA/booking-service/internal/usecase"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"go.uber.org/zap"
)

type BookingHandler struct {
	usecase *usecase.BookingUsecase
}

func NewBookingHandler(uc *usecase.BookingUsecase) *BookingHandler {
	return &BookingHandler{usecase: uc}
}

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		HouseID   string `json:"house_id"`
		UserID    string `json:"user_id"`
		DateStart string `json:"date_start"`
		DateEnd   string `json:"date_end"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Invalid JSON", zap.Error(err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(r.Context(), req.HouseID, req.UserID, req.DateStart, req.DateEnd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking created successfully"})
}
