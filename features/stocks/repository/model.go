package repository

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey"`
	BookID         uint      `gorm:"not null;index"`
	TotalStock     int       `gorm:"not null;check:total_stock >= 0"`
	AvailableStock int       `gorm:"not null;check:available_stock >= 0"`
}