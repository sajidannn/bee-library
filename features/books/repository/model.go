package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Author      string `gorm:"type:varchar(255);not null"`
	Publisher   string `gorm:"type:varchar(255);not null"`
	Stock       int    `gorm:"not null"`
	CoverImage  string `gorm:"type:text"`
}
