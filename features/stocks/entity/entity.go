package entity

import "time"

type Stock struct {
	ID             uint      `json:"id"`
	BookID         uint      `json:"book_id"`
	TotalStock     int       `json:"total_stock"`
	AvailableStock int       `json:"available_stock"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
