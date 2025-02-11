package handler

import (
	"bee-library/features/return_transactions/service"
	"bee-library/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReturnTransactionHandler struct {
	service service.ReturnTransactionService
}

func NewReturnTransactionHandler(service service.ReturnTransactionService) *ReturnTransactionHandler {
	return &ReturnTransactionHandler{service: service}
}

func (h *ReturnTransactionHandler) CreateReturnTransaction(c *gin.Context) {
	var req CreateReturnTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	returnTransaction, err := h.service.CreateReturnTransaction(req.BorrowTransactionID, req.ReturnDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helper.Response{
		Status:  "success",
		Message: "Return transaction created successfully",
		Data:    ToReturnTransactionResponse(*returnTransaction),
	})
}

func (h *ReturnTransactionHandler) GetAllReturnTransactions(c *gin.Context) {
	returnTransactions, err := h.service.GetAllReturnTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError{
			Status:  "error",
			Message: "Failed to fetch return transactions",
		})
		return
	}

	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Return transactions retrieved successfully",
		Data:    ToReturnTransactionResponseList(returnTransactions),
	})
}

func (h *ReturnTransactionHandler) GetReturnTransactionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: "Invalid return transaction ID",
		})
		return
	}

	returnTransaction, err := h.service.GetReturnTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Return transaction retrieved successfully",
		Data:    ToReturnTransactionResponse(*returnTransaction),
	})
}
