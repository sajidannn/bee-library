package repository

import (
	"bee-library/features/books/entity"
	"errors"

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
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*entity.Book, error) {
	var book entity.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bookRepository) Create(book *entity.Book) error {
	return r.db.Create(book).Error
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