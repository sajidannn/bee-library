package entity

import "time"

type ReturnReport struct {
	ID                  uint      `json:"id"`
	BookID              uint      `json:"book_id"`
	BorrowTransactionID uint      `json:"borrow_transaction_id"`
	ReturnTransactionID uint      `json:"return_transaction_id"`
	CreatedAt           time.Time `json:"created_at"`	
}

type ReturnReports struct {
	ReturnReport
	BookName     string `json:"book_name"`
	BorrowerName string `json:"borrower_name"`
}

type ReturnReportDetail struct {
	ReturnReport
	BookName         string    `json:"book_name"`
	BorrowerName     string    `json:"borrower_name"`
	ReturnDate       time.Time `json:"return_date"`
	FineAmount       float64   `json:"fine_amount"`
}

type ReturnReportRepository interface {
	GetAllReports(bookID *uint, memberID *uint, startDate, endDate *time.Time) ([]ReturnReports, error)
	GetReportByID(id uint) (*ReturnReportDetail, error)
}

type ReturnReportService interface {
	GetAllReports(bookID *uint, memberID *uint, startDate, endDate *time.Time) ([]ReturnReports, error)
	GetReportByID(id uint) (*ReturnReportDetail, error)
}