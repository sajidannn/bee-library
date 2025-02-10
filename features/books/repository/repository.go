package repository

import (
	"bee-library/features/books/entity"
	entityStock "bee-library/features/stocks/entity"
	"bee-library/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type bookWithStock struct {
	entity.Book     
	TotalStock     int `json:"total_stock"`
	AvailableStock int `json:"available_stock"`
}

func NewBookRepository(db *gorm.DB) entity.BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Select("id, title, author, publisher, category").
		Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*entity.Book, int, int, error) {
	var bookData bookWithStock

	err := r.db.Raw(`
		SELECT b.id, b.title, b.author, b.publisher, b.category, b.isbn, b.year, b.cover_image, 
		       b.created_at, b.updated_at, 
		       COALESCE(s.total_stock, 0) AS total_stock, 
		       COALESCE(s.available_stock, 0) AS available_stock
		FROM books b
		LEFT JOIN stocks s ON b.id = s.book_id
		WHERE b.id = ?
	`, id).Scan(&bookData).Error

	if err != nil {
		return nil, 0, 0, err
	}

	if bookData.Book.ID == 0 {
		return nil, 0, 0, helper.ErrNotFound
	}

	return &bookData.Book, bookData.TotalStock, bookData.AvailableStock, nil
}

func (r *bookRepository) Create(book *entity.Book) error {
	tx := r.db.Begin()

	if err := tx.Create(book).Error; err != nil {
		tx.Rollback()
		return err
	}

	stock := entityStock.Stock{
		BookID:         book.ID,
		TotalStock:     1,
		AvailableStock: 1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := tx.Create(&stock).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *bookRepository) Update(id uint, updatedBook *entity.Book) error {
	return r.db.Model(&entity.Book{}).Where("id = ?", id).Updates(updatedBook).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Book{}, id).Error
}

func (r *bookRepository) IsIsbnExist(isbn string) (bool, error) {
	var book entity.Book
	err := r.db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *bookRepository) IsBookExist(id uint) (bool, error) {
	var book entity.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}