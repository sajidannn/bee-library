package service

import (
	"bee-library/features/return_reports/entity"
	"time"
)

type returnReportService struct {
	repo entity.ReturnReportRepository
}

func NewReturnReportService(repo entity.ReturnReportRepository) entity.ReturnReportService {
	return &returnReportService{repo: repo}
}

func (s *returnReportService) GetAllReports(bookID *uint, memberID *uint, startDate, endDate *time.Time) ([]entity.ReturnReports, error) {
	return s.repo.GetAllReports(bookID, memberID, startDate, endDate)
}

func (s *returnReportService) GetReportByID(id uint) (*entity.ReturnReportDetail, error) {
	return s.repo.GetReportByID(id)
}
