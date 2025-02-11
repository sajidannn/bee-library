package routes

import (
	"bee-library/db"
	"bee-library/features/return_transactions/handler"
	"bee-library/features/return_transactions/repository"
	"bee-library/features/return_transactions/service"

	"github.com/gin-gonic/gin"
)

func RegisterReturnTransactionRoutes(router *gin.Engine) {
	returnTransactionRepo := repository.NewReturnTransactionRepo(db.DB)
	returnTransactionservice := service.NewReturnTransactionService(returnTransactionRepo)
	returnTransactionhandler := handler.NewReturnTransactionHandler(returnTransactionservice)

	returnTransactions := router.Group("/return-transactions")
	{
		returnTransactions.GET("/", returnTransactionhandler.GetAllReturnTransactions)
		returnTransactions.GET("/:id", returnTransactionhandler.GetReturnTransactionByID)
		returnTransactions.POST("/", returnTransactionhandler.CreateReturnTransaction)
	}
}
