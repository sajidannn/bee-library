package returnreports

import (
	"bee-library/db"
	"bee-library/features/return_reports/handler"
	"bee-library/features/return_reports/repository"
	"bee-library/features/return_reports/service"

	"github.com/gin-gonic/gin"
)

func RegisterReturnReportRoutes(r *gin.Engine) {
	returnReportRepo := repository.NewReturnReportRepo(db.DB)
	returnReportService := service.NewReturnReportService(returnReportRepo)
	returnReportHandler := handler.NewReturnReportHandler(returnReportService)

	returnReportRoutes := r.Group("/return-reports")
	{
		returnReportRoutes.GET("/", returnReportHandler.GetAllReports)
		returnReportRoutes.GET("/:id", returnReportHandler.GetReportByID)
	}
}