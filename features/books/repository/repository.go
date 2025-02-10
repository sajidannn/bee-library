package repository

import (
	"bee-library/features/books/entity"
	entityStock "bee-library/features/stocks/entity"
	"errors"
	"time"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]entity.Book, error)
	GetByID(id uint) (*entity.Book, error)
	Create(book *entity.Book) error
	Update(id uint, updatedBook *entity.Book) error
	Delete(id uint) error
	IsIsbnExist(isbn string) (bool, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Select("id, title, author, publisher, category").
		Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*entity.Book, error) {
	var book entity.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
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