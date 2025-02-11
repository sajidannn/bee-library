package entity

import "time"

type ReturnReport struct {
	ID                  uint      `json:"id"`
	BookID              uint      `json:"book_id"`
	BorrowTransactionID uint      `json:"borrow_transaction_id"`
	ReturnTransactionID uint      `json:"return_transaction_id"`
	CreatedAt           time.Time `json:"created_at"`	
}