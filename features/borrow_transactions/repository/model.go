package repository

import (
	bookEntity "bee-library/features/books/entity"
	memberEntity "bee-library/features/members/entity"
	"time"

	"gorm.io/gorm"
)

type BorrowTransaction struct {
	gorm.Model
	ID         	uint      `gorm:"primaryKey"`
	MemberID   	uint      `gorm:"not null"`
	BookID     	uint      `gorm:"not null"`
	BorrowDate 	time.Time `gorm:"not null"`
	DueDate   	time.Time `gorm:"not null"`
	Status    	string    `gorm:"type:varchar(20);check:status IN ('borrowed', 'returned')"`
	Member     	memberEntity.Member  `gorm:"constraint:OnDelete:CASCADE;"`
	Book				bookEntity.Book  `gorm:"constraint:OnDelete:CASCADE;"`
}