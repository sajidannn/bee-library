package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Author      string `gorm:"type:varchar(255);not null"`
	Publisher   string `gorm:"type:varchar(255);"`
	Category    string `gorm:"type:varchar(255);"`
	Isbn				string `gorm:"type:varchar(255);UNIQUE;"`
	Year        string `gorm:"type:varchar(255);"`
	CoverImage  string `gorm:"type:text"`
}
