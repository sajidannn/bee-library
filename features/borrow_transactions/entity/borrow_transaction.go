package entity

import "time"

type BorrowTransaction struct {
	ID            uint      `json:"id"`
	MemberID      uint      `json:"member_id"`
	BookID				uint      `json:"book_id"`
	BorrowDate		time.Time `json:"borrow_date"`
	DueDate			 	time.Time `json:"due_date"`
	Status			  string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BorrowTransactionRepository interface {
	GetAll() ([]BorrowTransaction, error)
	GetByID(id uint) (*BorrowTransaction, error)
	Create(transaction *BorrowTransaction) error
}

type BorrowTransactionService interface {
	GetAllTransactions() ([]BorrowTransaction, error)
	GetTransactionByID(id uint) (*BorrowTransaction, error)
	CreateTransaction(transaction *BorrowTransaction) (*BorrowTransaction, error)
}