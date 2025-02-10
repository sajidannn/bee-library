package service

import (
	"bee-library/features/borrow_transactions/entity"
	"time"
)

type borrowTransactionService struct {
	repo entity.BorrowTransactionRepository
}

func NewBorrowTransactionService(repo entity.BorrowTransactionRepository) entity.BorrowTransactionService {
	return &borrowTransactionService{repo: repo}
}

func (s *borrowTransactionService) GetAllTransactions() ([]entity.BorrowTransaction, error) {
	return s.repo.GetAll()
}

func (s *borrowTransactionService) GetTransactionByID(id uint) (*entity.BorrowTransaction, error) {
	return s.repo.GetByID(id)
}

func (s *borrowTransactionService) CreateTransaction(newTransaction *entity.BorrowTransaction) (*entity.BorrowTransaction, error) {
	newTransaction.BorrowDate = time.Now()
	newTransaction.Status = "borrowed"
	err := s.repo.Create(newTransaction)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}
