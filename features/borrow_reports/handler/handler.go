package handler

import (
	"bee-library/features/borrow_reports/entity"
	"bee-library/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BorrowReportHandler struct {
	service entity.BorrowReportService
}

func NewBorrowReportHandler(service entity.BorrowReportService) *BorrowReportHandler {
	return &BorrowReportHandler{service: service}
}

func (h *BorrowReportHandler) GetAllReports(c *gin.Context) {
	var req GetBorrowReportsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: "Invalid request format",
		})
		return
	}

	var startDate, endDate *time.Time
	if req.StartDate != nil && req.EndDate != nil {
		start, err1 := time.Parse("2006-01-02", *req.StartDate)
		end, err2 := time.Parse("2006-01-02", *req.EndDate)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, helper.ResponseError{
				Status:  "error",
				Message: "Invalid date format, use YYYY-MM-DD",
			})
			return
		}
		startDate, endDate = &start, &end
	}

	reports, err := h.service.GetAllReports(req.BookID, startDate, endDate)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if len(reports) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "There's no data"})
		return
	}

	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Reports retrieved successfully",
		Data:    ToBorrowReportResponseList(reports),
	})
}

func (h *BorrowReportHandler) GetReportByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: "Invalid report ID",
		})
		return
	}

	report, err := h.service.GetReportByID(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Report retrieved successfully",
		Data:    ToBorrowReportDetailResponse(report),
	})
}

func (h *BorrowReportHandler) GetTotalBorrowCount(c *gin.Context) {
	bookIDParam := c.Param("book_id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: "Invalid report ID",
		})
		return
	}

	totalCount, err := h.service.GetTotalBorrowCount(uint(bookID))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book_id":      bookID,
		"total_borrow": totalCount,
	})
}