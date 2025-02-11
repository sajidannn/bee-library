package entity

import "time"

type ReturnTransaction struct {
	ID               		uint `json:"id"`
	BorrowTransactionID uint `json:"borrow_transaction_id"`
	ReturnDate					time.Time `json:"return_date"`
	FineAmount					float64 `json:"fine_amount"`
	CreatedAt						time.Time `json:"created_at"`
}


type ReturnTransactionRepository interface {
	Create(returnTransaction *ReturnTransaction) error
	GetAll() ([]ReturnTransaction, error)
	GetByID(id uint) (*ReturnTransaction, error)
}