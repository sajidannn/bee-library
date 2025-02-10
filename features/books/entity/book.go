package entity

import "time"

type Book struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Category   string    `json:"category"`
	Isbn       string    `json:"isbn"`
	Year       string    `json:"year"`
	CoverImage string    `json:"cover_image"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BookService interface {
	GetAllBooks() ([]Book, error)
	GetBookByID(id uint) (*Book, int, int, error)
	CreateBook(book *Book) error
	UpdateBook(id uint, updatedBook *Book) error
	DeleteBook(id uint) error
}

type BookRepository interface {
	GetAll() ([]Book, error)
	GetByID(id uint) (*Book, int, int, error)
	Create(book *Book) error
	Update(id uint, updatedBook *Book) error
	Delete(id uint) error
	IsIsbnExist(isbn string) (bool, error)
	IsBookExist(id uint) (bool, error)
}