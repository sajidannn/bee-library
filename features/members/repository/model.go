package repository

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name    	string `gorm:"type:varchar(255);not null"`
	Email   	string `gorm:"type:varchar(255);unique;not null"`
	Password 	string `gorm:"type:varchar(255);not null"`
	Phone   	string `gorm:"type:varchar(20);not null"`
	Address 	string `gorm:"type:text;not null"`
	Photo   	string `gorm:"type:text"`
	CreatedAt	 	string `gorm:"type:timestamp"`
	UpdatedAt	 	string `gorm:"type:timestamp"`
}
