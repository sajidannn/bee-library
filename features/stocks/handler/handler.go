package handler

import (
	"bee-library/features/stocks/entity"
	"bee-library/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	service entity.StockService
}

func NewStockHandler(service entity.StockService) *StockHandler {
	return &StockHandler{service: service}
}

func (h *StockHandler) GetAllStock(c *gin.Context) {
	stocks, err := h.service.GetAllStock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError{
			Status:  "error",
			Message: "Failed to fetch stocks",
		})
		return
	}
	if len(stocks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "There's no data"})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    ToStockResponseList(stocks),
	})
}

func (h *StockHandler) GetStockByBookID(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	stock, err := h.service.GetStockByBookID(uint(bookID))
	if err != nil || stock == nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status: "succes",
		Message: "Book retrieved successfully",
		Data:  ToStockResponse(*stock),
	})
}

func (h *StockHandler) UpdateStock(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	var req StockUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	updatedStock := entity.Stock{
		TotalStock:     req.TotalStock,
		AvailableStock: req.AvailableStock,
	}
	updatedStock.UpdatedAt = time.Now()

	err := h.service.UpdateStock(uint(bookID), &updatedStock)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Book updated successfully",
	})
}
