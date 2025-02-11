package repository

import (
	"bee-library/features/return_reports/entity"
	"bee-library/helper"
	"time"

	"gorm.io/gorm"
)

type returnReportRepo struct {
	db *gorm.DB
}

func NewReturnReportRepo(db *gorm.DB) entity.ReturnReportRepository {
	return &returnReportRepo{db: db}
}

func (r *returnReportRepo) GetAllReports(bookID *uint, memberID *uint, startDate, endDate *time.Time) ([]entity.ReturnReports, error) {
	var reports []entity.ReturnReports
	query := r.db.Table("return_reports").
		Select("return_reports.id, return_reports.created_at, return_reports.updated_at, return_reports.deleted_at, "+
			"return_reports.return_transaction_id, return_reports.borrow_transaction_id, return_reports.book_id, "+
			"books.title AS book_name, members.name AS borrower_name, return_transactions.return_date").
		Joins("JOIN return_transactions ON return_reports.return_transaction_id = return_transactions.id").
		Joins("JOIN borrow_transactions ON return_reports.borrow_transaction_id = borrow_transactions.id").
		Joins("JOIN books ON return_reports.book_id = books.id").
		Joins("JOIN members ON borrow_transactions.member_id = members.id")

	if bookID != nil {
		query = query.Where("return_reports.book_id = ?", *bookID)
	}
	if memberID != nil {
		query = query.Where("borrow_transactions.member_id = ?", *memberID)
	}
	if startDate != nil && endDate != nil {
		query = query.Where("return_reports.created_at BETWEEN ? AND ?", *startDate, *endDate)
	}

	err := query.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *returnReportRepo) GetReportByID(id uint) (*entity.ReturnReportDetail, error) {
	var report entity.ReturnReportDetail
	err := r.db.Table("return_reports").
		Select("return_reports.id, return_reports.created_at, return_reports.return_transaction_id, "+
			"return_reports.borrow_transaction_id, return_reports.book_id, books.title AS book_name, "+
			"members.name AS borrower_name, return_transactions.return_date, return_transactions.fine_amount").
		Joins("JOIN return_transactions ON return_reports.return_transaction_id = return_transactions.id").
		Joins("JOIN borrow_transactions ON return_reports.borrow_transaction_id = borrow_transactions.id").
		Joins("JOIN books ON return_reports.book_id = books.id").
		Joins("JOIN members ON borrow_transactions.member_id = members.id").
		Where("return_reports.id = ?", id).
		First(&report).Error

	if err != nil {
		return nil, helper.ErrNotFound
	}
	
	return &report, nil
}

