package domain

type House struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id,omitempty"`
	HostelID      string `json:"hostel_id"`
	Address       string `json:"address"`
	NumberOfRooms int    `json:"number_of_rooms"`
	PricePerDay   int    `json:"price_per_day"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
}