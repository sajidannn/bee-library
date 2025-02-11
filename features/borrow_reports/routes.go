package borrowreports

import (
	"bee-library/db"
	"bee-library/features/borrow_reports/handler"
	"bee-library/features/borrow_reports/repository"
	"bee-library/features/borrow_reports/service"

	"github.com/gin-gonic/gin"
)

func RegisterBorrowReportRoutes(r *gin.Engine) {
	borrowReportRepo := repository.NewBorrowReportRepo(db.DB)
	borrowReportService := service.NewBorrowReportService(borrowReportRepo)
	borrowReportHandler := handler.NewBorrowReportHandler(borrowReportService)

	borrowReportRoutes := r.Group("/borrow-reports")
	{
		borrowReportRoutes.GET("/", borrowReportHandler.GetAllReports)
		borrowReportRoutes.GET("/:id", borrowReportHandler.GetReportByID)
		borrowReportRoutes.GET("/count/:book_id", borrowReportHandler.GetTotalBorrowCount)
	}
}