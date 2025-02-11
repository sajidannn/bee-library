package main

import (
	"bee-library/db"
	"bee-library/features/books"
	bookRepo "bee-library/features/books/repository"
	borrowReport "bee-library/features/borrow_reports"
	borrowReportRepo "bee-library/features/borrow_reports/repository"
	borrowTransaction "bee-library/features/borrow_transactions"
	borrowTransactionRepo "bee-library/features/borrow_transactions/repository"
	"bee-library/features/members"
	memberRepo "bee-library/features/members/repository"
	returnReport "bee-library/features/return_reports"
	returnReportRepo "bee-library/features/return_reports/repository"
	returnTransaction "bee-library/features/return_transactions"
	returnTransactiontRepo "bee-library/features/return_transactions/repository"
	"bee-library/features/stocks"
	stockRepo "bee-library/features/stocks/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	err := db.DB.AutoMigrate(
		&bookRepo.Book{},
		&memberRepo.Member{},
		&stockRepo.Stock{},
		&borrowTransactionRepo.BorrowTransaction{},
		&borrowReportRepo.BorrowReport{},
		&returnTransactiontRepo.ReturnTransaction{},
		&returnReportRepo.ReturnReport{},
	)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from bee library!"})
	})
	members.RegisterMemberRoutes(r)
	books.RegisterBookRoutes(r)
	stocks.RegisterStockRoutes(r)
	borrowTransaction.RegisterBorrowTransactionRoutes(r)
	borrowReport.RegisterBorrowReportRoutes(r)
	returnTransaction.RegisterReturnTransactionRoutes(r)
	returnReport.RegisterReturnReportRoutes(r)

	r.Run(":8080")
}
