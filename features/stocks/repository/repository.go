package repository

import (
	"bee-library/features/stocks/entity"
	"bee-library/helper"
	"errors"

	"gorm.io/gorm"
)

type StockRepository interface {
	GetAll() ([]entity.Stock, error)
	GetByID(bookID uint) (*entity.Stock, error)
	Update(bookID uint, updatedStock entity.Stock) error
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) GetAll() ([]entity.Stock, error) {
	var stocks []entity.Stock
	if err := r.db.Find(&stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}

func (r *stockRepository) GetByID(bookID uint) (*entity.Stock, error) {
	var stock entity.Stock
	err := r.db.First(&stock, bookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrNotFound
		}
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) Update(bookID uint, updatedStock entity.Stock) error {
	return r.db.Model(&entity.Stock{}).Where("book_id = ?", bookID).Updates(updatedStock).Error
}
