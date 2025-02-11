package handler

import "time"

// Request untuk membuat transaksi pengembalian
type CreateReturnTransactionRequest struct {
	BorrowTransactionID uint      `json:"borrow_transaction_id" binding:"required"`
	ReturnDate          time.Time `json:"return_date" binding:"required"`
}
