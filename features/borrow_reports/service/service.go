package service

import (
	"bee-library/features/borrow_reports/entity"
	"time"
)

type borrowReportService struct {
	repo entity.BorrowReportRepository
}

func NewBorrowReportService(repo entity.BorrowReportRepository) entity.BorrowReportService {
	return &borrowReportService{repo: repo}
}

func (s *borrowReportService) GetAllReports(bookID *uint, startDate, endDate *time.Time) ([]entity.BorrowReports, error) {
	reports, err := s.repo.GetAllReports(bookID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (s *borrowReportService) GetReportByID(id uint) (*entity.BorrowReportDetail, error) {
	report, err := s.repo.GetReportByID(id)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (s *borrowReportService) GetTotalBorrowCount(bookID uint) (int64, error) {
	count, err := s.repo.GetTotalBorrowCount(bookID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
