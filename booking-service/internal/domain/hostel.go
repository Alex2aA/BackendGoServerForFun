package domain

import "time"

type Hostel struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Rate        int       `json:"rate"` // 1-5
	CreatedAt   time.Time `json:"created_at"`
}