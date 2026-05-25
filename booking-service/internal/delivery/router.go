package http

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(
	hostelH *HostelHandler,
	houseH *HouseHandler,
	bookingH *BookingHandler,
) *mux.Router {

	r := mux.NewRouter()

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api").Subrouter()

	// Hostels
	api.HandleFunc("/hostel", hostelH.Create).Methods("POST")

	// Houses
	api.HandleFunc("/house", houseH.Create).Methods("POST")

	// Bookings
	api.HandleFunc("/booking", bookingH.Create).Methods("POST")

	return r
}