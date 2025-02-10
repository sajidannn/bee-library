package entity

import "time"

type Book struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Category   string    `json:"category"`
	Isbn       string    `json:"isbn"`
	Year       string    `json:"year"`
	CoverImage string    `json:"cover_image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
