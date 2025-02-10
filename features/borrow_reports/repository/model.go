package repository

import (
	bookEntity "bee-library/features/books/entity"
	borrowTransactionEntity "bee-library/features/borrow_transactions/entity"

	"time"
)

type BorrowReport struct {
	ID         						uint `gorm:"primaryKey"`
	BorrowTransactionID   uint `gorm:"not null"`
	BookID     						uint `gorm:"not null"`
	CreatedAt  						time.Time
	BorrowTransaction 		borrowTransactionEntity.BorrowTransaction `gorm:"constraint:OnDelete:CASCADE;"`
	Book									bookEntity.Book  `gorm:"constraint:OnDelete:CASCADE;"`
}