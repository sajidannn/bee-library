package repository

import (
	"bee-library/features/borrow_transactions/entity"

	"gorm.io/gorm"
)

type ReturnTransaction struct {
	gorm.Model
	ID               		uint `gorm:"primaryKey"`
	BorrowTransactionID uint `gorm:"not null"`
	ReturnDate					string `gorm:"type:timestamp"`
	FineAmount					float64 `gorm:"not null"`
	CreatedAt						string `gorm:"type:timestamp"`
	BorrowTransaction		entity.BorrowTransaction `gorm:"constraint:OnDelete:CASCADE;"`
}