package repository

import (
	borrowReportEntity "bee-library/features/borrow_reports/entity"
	borrowTransactionEntity "bee-library/features/borrow_transactions/entity"
	memberEntity "bee-library/features/members/entity"
	stockEntity "bee-library/features/stocks/entity"
	"bee-library/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type borrowTransactionRepo struct {
	db *gorm.DB
}

func NewBorrowTransactionRepository(db *gorm.DB) borrowTransactionEntity.BorrowTransactionRepository {
	return &borrowTransactionRepo{db: db}
}

func (r *borrowTransactionRepo) GetAll() ([]borrowTransactionEntity.BorrowTransaction, error) {
	var transactions []borrowTransactionEntity.BorrowTransaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *borrowTransactionRepo) GetByID(id uint) (*borrowTransactionEntity.BorrowTransaction, error) {
	var transaction borrowTransactionEntity.BorrowTransaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrNotFound
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *borrowTransactionRepo) Create(transaction *borrowTransactionEntity.BorrowTransaction) error {
	var stock stockEntity.Stock
	var member memberEntity.Member

	err := r.db.First(&member, transaction.MemberID).Error
	if err != nil {
		return errors.New("member not found")
	}
	err = r.db.Where("book_id = ?", transaction.BookID).First(&stock).Error
	if err != nil {
		return errors.New("book not found")
	}
	if stock.AvailableStock <= 0 {
		return errors.New("no stock available")
	}

	if err := r.db.Create(transaction).Error; err != nil {
		return err
	}

	stock.AvailableStock -= 1
	if err := r.db.Save(&stock).Error; err != nil {
		return err
	}

	report := borrowReportEntity.BorrowReport{
		BorrowTransactionID: 	transaction.ID,
		BookID:   						transaction.BookID,
		CreatedAt: 						time.Now(),
	}
	return r.db.Create(&report).Error
}
