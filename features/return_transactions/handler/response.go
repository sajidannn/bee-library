package handler

import (
	"bee-library/features/return_transactions/entity"
	"time"
)

type ReturnTransactionResponse struct {
	ID                   uint      `json:"id"`
	BorrowTransactionID  uint      `json:"borrow_transaction_id"`
	ReturnDate           time.Time `json:"return_date"`
	FineAmount           float64   `json:"fine_amount"`
	CreatedAt            time.Time `json:"created_at"`
}

func ToReturnTransactionResponse(tx entity.ReturnTransaction) ReturnTransactionResponse {
	return ReturnTransactionResponse{
		ID:                   tx.ID,
		BorrowTransactionID:  tx.BorrowTransactionID,
		ReturnDate:           tx.ReturnDate,
		FineAmount:           tx.FineAmount,
		CreatedAt:            tx.CreatedAt,
	}
}

func ToReturnTransactionResponseList(txs []entity.ReturnTransaction) []ReturnTransactionResponse {
	var responseList []ReturnTransactionResponse
	for _, tx := range txs {
		responseList = append(responseList, ToReturnTransactionResponse(tx))
	}
	return responseList
}
