package service

import (
	"bee-library/features/books/entity"
	"bee-library/features/books/repository"
	"bee-library/helper"
)

type BookService interface {
	GetAllBooks() ([]entity.Book, error)
	GetBookByID(id uint) (*entity.Book, error)
	CreateBook(book *entity.Book) error
	UpdateBook(id uint, updatedBook *entity.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]entity.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBookByID(id uint) (*entity.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, helper.ErrNotFound
	}
	return book, nil
}

func (s *bookService) CreateBook(newBook *entity.Book) error {
	exists, err := s.repo.IsIsbnExist(newBook.Isbn)
	if err != nil {
		return helper.ErrInternalServer
	}
	if exists {
		return helper.ErrIsbnExists
	}
	if err := s.repo.Create(newBook); err != nil {
		return helper.ErrInternalServer
	}

	return nil
}

func (s *bookService) UpdateBook(id uint, updatedBook *entity.Book) error {
	_, err := s.GetBookByID(id)
	if err != nil {
		return err
	}
	if err := s.repo.Update(id, updatedBook); err != nil {
		return helper.ErrInternalServer
	}
	return nil
}

func (s *bookService) DeleteBook(id uint) error {
	_, err := s.GetBookByID(id)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(id); err != nil {
		return helper.ErrInternalServer
	}
	return nil
}