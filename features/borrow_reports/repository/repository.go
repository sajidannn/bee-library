package repository

import (
	"time"

	"bee-library/features/borrow_reports/entity"
	"bee-library/helper"

	"gorm.io/gorm"
)

type borrowReportRepo struct {
	db *gorm.DB
}

func NewBorrowReportRepo(db *gorm.DB) entity.BorrowReportRepository {
	return &borrowReportRepo{db: db}
}

func (r *borrowReportRepo) GetAllReports(bookID *uint, startDate, endDate *time.Time) ([]entity.BorrowReports, error) {
	var reports []entity.BorrowReports

	query := r.db.Table("borrow_reports").
	Select("borrow_reports.*, books.title AS book_name, members.name AS borrower_name").
	Joins("JOIN borrow_transactions ON borrow_reports.borrow_transaction_id = borrow_transactions.id").
	Joins("JOIN books ON borrow_transactions.book_id = books.id").
	Joins("JOIN members ON borrow_transactions.member_id = members.id")

if bookID != nil {
	query = query.Where("books.id = ?", *bookID)
}
if startDate != nil && endDate != nil {
	query = query.Where("borrow_reports.created_at BETWEEN ? AND ?", *startDate, *endDate)
}
if startDate != nil && endDate == nil {
	query = query.Where("borrow_reports.created_at >= ?", *startDate)
}
if endDate != nil && startDate == nil {
	query = query.Where("borrow_reports.created_at <= ?", *endDate)
}

	err := query.Find(&reports).Error
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *borrowReportRepo) GetReportByID(id uint) (*entity.BorrowReportDetail, error) {
	var report entity.BorrowReportDetail

	err := r.db.Table("borrow_reports").
		Select(`borrow_reports.*, books.title AS book_name, 
		        members.name AS borrower_name,
		        borrow_transactions.borrow_date, borrow_transactions.status, borrow_transactions.due_date`).
		Joins("JOIN borrow_transactions ON borrow_reports.borrow_transaction_id = borrow_transactions.id").
		Joins("JOIN books ON borrow_transactions.book_id = books.id").
		Joins("JOIN members ON borrow_transactions.member_id = members.id").
		Where("borrow_reports.id = ?", id).
		First(&report).Error

	if err != nil {
		return nil, helper.ErrNotFound
	}

	return &report, nil
}

func (r *borrowReportRepo) GetTotalBorrowCount(bookID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entity.BorrowReport{}).
		Joins("JOIN borrow_transactions ON borrow_reports.borrow_transaction_id = borrow_transactions.id").
		Where("borrow_transactions.book_id = ?", bookID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
