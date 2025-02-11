package handler

import (
	"bee-library/features/borrow_reports/entity"
	"time"
)

type BorrowReportResponse struct {
	ID            uint      `json:"id"`
	BorrowID      uint      `json:"borrow_transaction_id"`
	BookID        uint      `json:"book_id"`
	BookName      string    `json:"book_name"`
	BorrowerName  string    `json:"borrower_name"`
	CreatedAt     time.Time `json:"created_at"`
}

type BorrowReportDetailResponse struct {
	ID            uint      `json:"id"`
	BorrowID      uint      `json:"borrow_transaction_id"`
	BookID        uint      `json:"book_id"`
	BookName      string    `json:"book_name"`
	BorrowerName  string    `json:"borrower_name"`
	BorrowDate    time.Time `json:"borrow_date"`
	DueDate       time.Time `json:"due_date"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

func ToBorrowReportResponse(report entity.BorrowReports) BorrowReportResponse {
	return BorrowReportResponse{
		ID:            report.ID,
		BorrowID:      report.BorrowTransactionID,
		BookID:        report.BookID,
		BookName:      report.BookName,
		BorrowerName:  report.BorrowerName,
		CreatedAt:     report.CreatedAt,
	}
}

func ToBorrowReportDetailResponse(report *entity.BorrowReportDetail) BorrowReportDetailResponse {
	return BorrowReportDetailResponse{
		ID:            report.ID,
		BorrowID:      report.BorrowTransactionID,
		BookID:        report.BookID,
		BookName:      report.BookName,
		BorrowerName:  report.BorrowerName,
		BorrowDate:    report.BorrowDate,
		DueDate:       report.DueDate,
		Status:        report.Status,
		CreatedAt:     report.CreatedAt,
	}
}

func ToBorrowReportResponseList(reports []entity.BorrowReports) []BorrowReportResponse {
	var responseList []BorrowReportResponse
	for _, report := range reports {
		responseList = append(responseList, ToBorrowReportResponse(report))
	}
	return responseList
}
