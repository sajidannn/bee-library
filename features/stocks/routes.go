package stocks

import (
	"bee-library/db"
	"bee-library/features/stocks/handler"
	"bee-library/features/stocks/repository"
	"bee-library/features/stocks/service"

	"github.com/gin-gonic/gin"
)

func RegisterStockRoutes(r *gin.Engine) {
	stockRepo := repository.NewStockRepository(db.DB)
	stockService := service.NewStockService(stockRepo)
	stockHandler := handler.NewStockHandler(stockService)

	stockRoutes := r.Group("/stocks")
	{
		stockRoutes.GET("/", stockHandler.GetAllStock)
		stockRoutes.GET("/:book_id", stockHandler.GetStockByBookID)
		stockRoutes.PUT("/:book_id", stockHandler.UpdateStock)
	}
}