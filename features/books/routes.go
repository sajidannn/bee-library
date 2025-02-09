package books

import (
	"bee-library/db"
	"bee-library/features/books/handler"
	"bee-library/features/books/repository"
	"bee-library/features/books/service"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine) {
	bookRepo := repository.NewBookRepository(db.DB)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	bookRoutes := r.Group("/books") 
	{
		bookRoutes.GET("/", bookHandler.GetAllBooks)
		bookRoutes.GET("/:id", bookHandler.GetBookByID)
		bookRoutes.POST("/", bookHandler.CreateBook)
		bookRoutes.PUT("/:id", bookHandler.UpdateBook)
		bookRoutes.DELETE("/:id", bookHandler.DeleteBook)
	}

}