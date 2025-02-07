package main

import (
	"bee-library/db"
	bookRepo "bee-library/features/books/repository"
	"bee-library/features/members"
	memberRepo "bee-library/features/members/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	err := db.DB.AutoMigrate(
		&bookRepo.Book{},
		&memberRepo.Member{},
	)
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}

	r := gin.Default()
	members.RegisterMemberRoutes(r)

	r.Run(":8000")
}
