package service

import (
	"bee-library/features/stocks/entity"
	"bee-library/helper"
	"errors"
)

type stockService struct {
	repo entity.StockRepository
}

func NewStockService(repo entity.StockRepository) entity.StockService {
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