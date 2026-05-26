package http

import (
	"encoding/json"
	"net/http"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/internal/usecase"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"go.uber.org/zap"
)

type HouseHandler struct {
	usecase *usecase.HouseUsecase
}

func NewHouseHandler(uc *usecase.HouseUsecase) *HouseHandler {
	return &HouseHandler{usecase: uc}
}

func (h *HouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var house domain.House
	if err := json.NewDecoder(r.Body).Decode(&house); err != nil {
		logger.Log.Error("Invalid JSON", zap.Error(err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(r.Context(), &house); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "House created successfully"})
}
