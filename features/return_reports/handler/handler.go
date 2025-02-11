package handler

import (
	"bee-library/features/return_reports/entity"
	"bee-library/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReturnReportHandler struct {
	service entity.ReturnReportService
}

func NewReturnReportHandler(service entity.ReturnReportService) *ReturnReportHandler {
	return &ReturnReportHandler{service: service}
}

func (h *ReturnReportHandler) GetAllReports(c *gin.Context) {
	var req GetReturnReportsRequest
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

	reports, err := h.service.GetAllReports(req.BookID, req.MemberID, startDate, endDate)
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
		Data:    ToReturnReportResponseList(reports),
	})
}

func (h *ReturnReportHandler) GetReportByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: "Invalid ID format",
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
		Data:    ToReturnReportDetailResponse(report),
	})
}
