package service

import (
	"bee-library/features/books/entity"
	"bee-library/helper"
)

type bookService struct {
	repo entity.BookRepository
}

func NewBookService(repo entity.BookRepository) entity.BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]entity.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBookByID(id uint) (*entity.Book, int, int, error) {
	book, totalStock, availableStock, err := s.repo.GetByID(id)
	if err != nil {
		return nil, 0, 0, err
	}
	return book, totalStock, availableStock, nil
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
	_, err := s.repo.IsBookExist(id)
	if err != nil {
		return err
	}
	if err := s.repo.Update(id, updatedBook); err != nil {
		return helper.ErrInternalServer
	}
	return nil
}

func (s *bookService) DeleteBook(id uint) error {
	_, err := s.repo.IsBookExist(id)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(id); err != nil {
		return helper.ErrInternalServer
	}
	return nil
}