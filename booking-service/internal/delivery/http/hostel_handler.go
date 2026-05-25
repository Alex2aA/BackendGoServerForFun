package http

import (
	"encoding/json"
	"net/http"

	"github.com/Alex2aA/booking-service/internal/domain"
	"github.com/Alex2aA/booking-service/internal/usecase"
	"github.com/Alex2aA/booking-service/pkg/logger"
	"go.uber.org/zap"
)

type HostelHandler struct {
	usecase *usecase.HostelUsecase
}

func NewHostelHandler(uc *usecase.HostelUsecase) *HostelHandler {
	return &HostelHandler{usecase: uc}
}

// CreateHostel godoc
// @Summary Create new hostel
// @Tags hostel
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param hostel body domain.Hostel true "Hostel data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/hostel [post]
func (h *HostelHandler) Create(w http.ResponseWriter, r *http.Request) {
	var hostel domain.Hostel
	if err := json.NewDecoder(r.Body).Decode(&hostel); err != nil {
		logger.Log.Error("Invalid JSON", zap.Error(err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(r.Context(), &hostel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Hostel created successfully"})
}