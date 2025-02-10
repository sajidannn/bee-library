package entity

import "time"

type Member struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Photo     string `json:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
