package handler

import (
	"bee-library/features/return_reports/entity"
	"time"
)

type ReturnReportResponse struct {
	ID            uint      `json:"id"`
	ReturnID      uint      `json:"return_transaction_id"`
	BookID        uint      `json:"book_id"`
	BookName      string    `json:"book_name"`
	BorrowerName  string    `json:"borrower_name"`
	CreatedAt     time.Time `json:"created_at"`
}

type ReturnReportDetailResponse struct {
	ID            uint      `json:"id"`
	ReturnID      uint      `json:"return_transaction_id"`
	BookID        uint      `json:"book_id"`
	BookName      string    `json:"book_name"`
	BorrowerName  string    `json:"borrower_name"`
	ReturnDate    time.Time `json:"return_date"`
	FineAmount    float64   `json:"fine_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func ToReturnReportResponse(report entity.ReturnReports) ReturnReportResponse {
	return ReturnReportResponse{
		ID:            report.ID,
		ReturnID:      report.ReturnTransactionID,
		BookID:        report.BookID,
		BookName:      report.BookName,
		BorrowerName:  report.BorrowerName,
		CreatedAt:     report.CreatedAt,
	}
}

func ToReturnReportDetailResponse(report *entity.ReturnReportDetail) ReturnReportDetailResponse {
	return ReturnReportDetailResponse{
		ID:            report.ID,
		ReturnID:      report.ReturnTransactionID,
		BookID:        report.BookID,
		BookName:      report.BookName,
		BorrowerName:  report.BorrowerName,
		ReturnDate:    report.ReturnDate,
		FineAmount:    report.FineAmount,
		CreatedAt:     report.CreatedAt,
	}
}

func ToReturnReportResponseList(reports []entity.ReturnReports) []ReturnReportResponse {
	var responseList []ReturnReportResponse
	for _, report := range reports {
		responseList = append(responseList, ToReturnReportResponse(report))
	}
	return responseList
}