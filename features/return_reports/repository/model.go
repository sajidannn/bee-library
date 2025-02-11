package repository

import (
	book "bee-library/features/books/entity"
	borrowTrans "bee-library/features/borrow_transactions/entity"
	returnTrans "bee-library/features/return_transactions/entity"

	"gorm.io/gorm"
)

type ReturnReport struct {
	gorm.Model
	ID                   uint   `gorm:"primaryKey"`
	ReturnTransactionID  uint   `gorm:"not null"`
	BorrowTransactionID   uint   `gorm:"not null"`
	BookID               uint   `gorm:"not null"`
	CreatedAt            string `gorm:"type:timestamp"`
	ReturnTransaction    returnTrans.ReturnTransaction `gorm:"constraint:OnDelete:CASCADE;"`
	BorrowTransaction    borrowTrans.BorrowTransaction `gorm:"constraint:OnDelete:CASCADE;"`
	Book                 book.Book `gorm:"constraint:OnDelete:CASCADE;foreignKey:BookID"`
}
