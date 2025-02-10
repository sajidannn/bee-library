package handler

import "time"

type BorrowTransactionRequest struct {
	MemberID uint      `json:"member_id" binding:"required"`
	BookID   uint      `json:"book_id" binding:"required"`
	DueDate  time.Time `json:"due_date" binding:"required"`
}
