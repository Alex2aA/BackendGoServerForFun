package domain

import "time"

type Booking struct {
	ID        string    `json:"id"`
	HouseID   string    `json:"house_id"`
	HostelID  string    `json:"hostel_id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	CreatedAt time.Time `json:"created_at"`
}