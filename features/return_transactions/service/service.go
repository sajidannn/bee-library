package service

import (
	"bee-library/features/return_transactions/entity"
	"time"
)

type ReturnTransactionService interface {
	CreateReturnTransaction(borrowTransactionID uint, returnDate time.Time) (*entity.ReturnTransaction, error)
	GetAllReturnTransactions() ([]entity.ReturnTransaction, error)
	GetReturnTransactionByID(id uint) (*entity.ReturnTransaction, error)
}

type returnTransactionService struct {
	repo entity.ReturnTransactionRepository
}

func NewReturnTransactionService(repo entity.ReturnTransactionRepository) ReturnTransactionService {
	return &returnTransactionService{repo: repo}
}

func (s *returnTransactionService) CreateReturnTransaction(borrowTransactionID uint, returnDate time.Time) (*entity.ReturnTransaction, error) {
	returnTransaction := &entity.ReturnTransaction{
		BorrowTransactionID: borrowTransactionID,
		ReturnDate:          returnDate,
		CreatedAt:           time.Now(),
	}

	err := s.repo.Create(returnTransaction)
	if err != nil {
		return nil, err
	}

	return returnTransaction, nil
}

func (s *returnTransactionService) GetAllReturnTransactions() ([]entity.ReturnTransaction, error) {
	return s.repo.GetAll()
}

func (s *returnTransactionService) GetReturnTransactionByID(id uint) (*entity.ReturnTransaction, error) {
	return s.repo.GetByID(id)
}
