package borrowtransactions

import (
	"bee-library/db"
	"bee-library/features/borrow_transactions/handler"
	"bee-library/features/borrow_transactions/repository"
	"bee-library/features/borrow_transactions/service"

	"github.com/gin-gonic/gin"
)

func RegisterBorrowTransactionRoutes(r *gin.Engine) {
	borrowTransactionRepo := repository.NewBorrowTransactionRepository(db.DB)
	borrowTransactionService := service.NewBorrowTransactionService(borrowTransactionRepo)
	borrowTransactionHandler := handler.NewBorrowTransactionHandler(borrowTransactionService)

	borrowTransactionRoutes := r.Group("/borrow_transactions")
	{
		borrowTransactionRoutes.GET("/", borrowTransactionHandler.GetAllTransactions)
		borrowTransactionRoutes.GET("/:id", borrowTransactionHandler.GetTransactionByID)
		borrowTransactionRoutes.POST("/", borrowTransactionHandler.CreateTransaction)
	}
}