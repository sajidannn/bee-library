package main

import (
	"bee-library/db"
	"bee-library/features/books"
	bookRepo "bee-library/features/books/repository"
	"bee-library/features/members"
	memberRepo "bee-library/features/members/repository"
	stockRepo "bee-library/features/stocks/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	err := db.DB.AutoMigrate(
		&bookRepo.Book{},
		&memberRepo.Member{},
		&stockRepo.Stock{},
	)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}

	r := gin.Default()
	members.RegisterMemberRoutes(r)
	books.RegisterBookRoutes(r)

	r.Run(":8000")
}
