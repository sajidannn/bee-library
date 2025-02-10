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

type StockRepository interface {
	GetAll() ([]Stock, error)
	GetByID(bookID uint) (*Stock, error)
	Update(bookID uint, updatedStock Stock) error
}

type StockService interface {
	GetAllStock() ([]Stock, error)
	GetStockByBookID(bookID uint) (*Stock, error)
	UpdateStock(bookID uint, updatedStock *Stock) error
}