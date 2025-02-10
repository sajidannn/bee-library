package entity

import "time"

type BorrowReport struct {
	ID                  uint      `json:"id"`
	BorrowTransactionID uint      `json:"borrow_transaction_id"`
	BookID              uint      `json:"book_id"`
	CreatedAt           time.Time `json:"created_at"`
}