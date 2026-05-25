package domain

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password,omitempty"`
	Alias        string `json:"alias"`
	RefreshToken string `json:"refresh_token,omitempty"`
}