package repository

import (
	borrowTransactionEntity "bee-library/features/borrow_transactions/entity"
	returnReportEntity "bee-library/features/return_reports/entity"
	returnTransactionEntity "bee-library/features/return_transactions/entity"
	stockEntity "bee-library/features/stocks/entity"
	"bee-library/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type returnTransactionRepo struct {
	db *gorm.DB
}

func NewReturnTransactionRepo(db *gorm.DB) returnTransactionEntity.ReturnTransactionRepository {
	return &returnTransactionRepo{db: db}
}

func (r *returnTransactionRepo) Create(returnTransaction *returnTransactionEntity.ReturnTransaction) error {
	var borrowTransaction borrowTransactionEntity.BorrowTransaction
	if err := r.db.Where("id = ? AND status = ?", returnTransaction.BorrowTransactionID, "borrowed").First(&borrowTransaction).Error; err != nil {
		return errors.New("borrow transaction not found or already returned")
	}

	fineAmount := 0.0
	if returnTransaction.ReturnDate.After(borrowTransaction.DueDate) {
		duration := returnTransaction.ReturnDate.Sub(borrowTransaction.DueDate).Hours()
		fineAmount = duration / 24 * 1000
	}

	returnTransaction.FineAmount = fineAmount

	if err := r.db.Create(returnTransaction).Error; err != nil {
		return err
	}

	if err := r.db.Model(&borrowTransaction).Update("status", "returned").Error; err != nil {
		return err
	}

	var stock stockEntity.Stock
	if err := r.db.Where("book_id = ?", borrowTransaction.BookID).First(&stock).Error; err != nil {
		return err
	}
	stock.AvailableStock += 1
	if err := r.db.Save(&stock).Error; err != nil {
		return err
	}

	report := returnReportEntity.ReturnReport{
		ReturnTransactionID: returnTransaction.ID,
		BorrowTransactionID: returnTransaction.BorrowTransactionID,
		BookID:              borrowTransaction.BookID,
		CreatedAt:           time.Now(),
	}
	if err := r.db.Create(&report).Error; err != nil {
		return err
	}

	return nil
}

func (r *returnTransactionRepo) GetAll() ([]returnTransactionEntity.ReturnTransaction, error) {
	var returnTransactions []returnTransactionEntity.ReturnTransaction
	err := r.db.Find(&returnTransactions).Error
	return returnTransactions, err
}

func (r *returnTransactionRepo) GetByID(id uint) (*returnTransactionEntity.ReturnTransaction, error) {
	var returnTransaction returnTransactionEntity.ReturnTransaction
	err := r.db.Where("id = ?", id).First(&returnTransaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrNotFound
		}
		return nil, err
	}
	return &returnTransaction, nil
}
