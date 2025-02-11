package entity

import "time"

type BorrowReport struct {
	ID                  uint      `json:"id"`
	BookID              uint      `json:"book_id"`
	BorrowTransactionID uint      `json:"borrow_transaction_id"`
	CreatedAt           time.Time `json:"created_at"`
}

type BorrowReports struct {
	BorrowReport
	BookName     string `json:"book_name"`
	BorrowerName string `json:"borrower_name"`
}
type BorrowReportDetail struct {
	BorrowReport
	BookName         string    `json:"book_name"`
	BorrowerName     string    `json:"borrower_name"`
	BorrowDate       time.Time `json:"borrow_date"`
	Status           string    `json:"status"`
	DueDate         time.Time `json:"due_date"`
}

type BorrowReportRepository interface {
	GetAllReports(bookID *uint, startDate, endDate *time.Time) ([]BorrowReports, error)
	GetReportByID(id uint) (*BorrowReportDetail, error)
	GetTotalBorrowCount(bookID uint) (int64, error)
}

type BorrowReportService interface {
	GetAllReports(bookID *uint, startDate, endDate *time.Time) ([]BorrowReports, error)
	GetReportByID(id uint) (*BorrowReportDetail, error)
	GetTotalBorrowCount(bookID uint) (int64, error)
}