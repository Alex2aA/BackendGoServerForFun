package http

import (
	"encoding/json"
	"net/http"

	"github.com/Alex2aA/booking-service/internal/usecase"
)

// @Summary Register
// @Tags user
// @Accept json
// @Produce json
// @Success 201
// @Router /api/user/register [post]
func RegisterHandler(uc *usecase.UserUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{ Username, Password string }
		json.NewDecoder(r.Body).Decode(&req)

		token, err := uc.Register(r.Context(), req.Username, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}