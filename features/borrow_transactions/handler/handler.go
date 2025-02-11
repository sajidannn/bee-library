package handler

import (
	"bee-library/features/borrow_transactions/entity"
	"bee-library/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BorrowTransactionHandler struct {
	service entity.BorrowTransactionService
}

func NewBorrowTransactionHandler(service entity.BorrowTransactionService) *BorrowTransactionHandler {
	return &BorrowTransactionHandler{service: service}
}

func (h *BorrowTransactionHandler) GetAllTransactions(c *gin.Context) {
	txs, err := h.service.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError{
			Status:  "error",
			Message: "Failed to fetch transactions",
		})
		return
	}
	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "There's no data"})
		return
	}
	c.JSON(http.StatusOK,helper.Response{
		Status:  "success",
		Message: "Transactions retrieved successfully",
		Data:    ToBorrowTransactionResponseList(txs),
	})
}

func (h *BorrowTransactionHandler) GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tx, err := h.service.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Transaction retrieved successfully",
		Data:    ToBorrowTransactionResponse(*tx),
	})
}

func (h *BorrowTransactionHandler) CreateTransaction(c *gin.Context) {
	var req BorrowTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	newTransaction := entity.BorrowTransaction{
		MemberID: req.MemberID,
		BookID:   req.BookID,
		DueDate:  req.DueDate,
	}

	tx, err := h.service.CreateTransaction(&newTransaction)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, helper.Response{
		Status:  "success",
		Message: "Transaction created successfully",
		Data:    ToBorrowTransactionResponse(*tx),
	})
}