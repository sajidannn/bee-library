package service

import (
	"bee-library/features/stocks/entity"
	"bee-library/features/stocks/repository"
	"bee-library/helper"
	"errors"
)

type StockService interface {
	GetAllStock() ([]entity.Stock, error)
	GetStockByBookID(bookID uint) (*entity.Stock, error)
	UpdateStock(bookID uint, updatedStock *entity.Stock) error
}

type stockService struct {
	repo repository.StockRepository
}

func NewStockService(repo repository.StockRepository) StockService {
	return &stockService{repo: repo}
}

func (s *stockService) GetAllStock() ([]entity.Stock, error) {
	return s.repo.GetAll()
}

func (s *stockService) GetStockByBookID(bookID uint) (*entity.Stock, error) {
	return s.repo.GetByID(bookID)
}

func (s *stockService) UpdateStock(bookID uint, updatedStock *entity.Stock) error {
	_, err := s.repo.GetByID(bookID)
	if err != nil {
		return err
	}

	if updatedStock.TotalStock < updatedStock.AvailableStock {
		return errors.New("available stock cannot be more than total stock")
	}

	if err := s.repo.Update(bookID, *updatedStock); err != nil {
		return helper.ErrInternalServer
	}

	return nil
}