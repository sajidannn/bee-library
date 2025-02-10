package handler

import (
	"bee-library/features/borrow_transactions/entity"
	"time"
)

type BorrowTransactionResponse struct {
	ID         uint      `json:"id"`
	MemberID   uint      `json:"member_id"`
	BookID     uint      `json:"book_id"`
	BorrowDate time.Time `json:"borrow_date"`
	DueDate    time.Time `json:"due_date"`
	Status     string    `json:"status"`
}

func ToBorrowTransactionResponse(tx entity.BorrowTransaction) BorrowTransactionResponse {
	return BorrowTransactionResponse{
		ID:         tx.ID,
		MemberID:   tx.MemberID,
		BookID:     tx.BookID,
		BorrowDate: tx.BorrowDate,
		DueDate:    tx.DueDate,
		Status:     tx.Status,
	}
}

func ToBorrowTransactionResponseList(txs []entity.BorrowTransaction) []BorrowTransactionResponse {
	var responseList []BorrowTransactionResponse
	for _, tx := range txs {
		responseList = append(responseList, ToBorrowTransactionResponse(tx))
	}
	return responseList
}
