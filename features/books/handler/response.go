package handler

import (
	"bee-library/features/books/entity"
)

type StockResponse struct {
	TotalStock     int `json:"total_stock"`
	AvailableStock int `json:"available_stock"`
}

type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Category  string `json:"category"`
}

// Response untuk detail buku (dengan informasi tambahan)
type BookDetailResponse struct {
	ID         uint          `json:"id"`
	Title      string        `json:"title"`
	Author     string        `json:"author"`
	Publisher  string        `json:"publisher"`
	Category   string        `json:"category"`
	Isbn       string        `json:"isbn"`
	Year       string        `json:"year"`
	CoverImage string        `json:"cover_image"`
	CreatedAt  string        `json:"created_at"`
	UpdatedAt  string        `json:"updated_at"`
	Stock      StockResponse `json:"stock"`
}

func ToBookResponse(book entity.Book) BookResponse {
	return BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author: 	book.Author,
		Publisher: book.Publisher,
		Category:  book.Category,
	}
}

func ToBookDetailResponse(book entity.Book, totalStock, availableStock int) BookDetailResponse {
	return BookDetailResponse{
		ID:         book.ID,
		Title:      book.Title,
		Author:     book.Author,
		Publisher:  book.Publisher,
		Category:   book.Category,
		Isbn:       book.Isbn,
		Year:       book.Year,
		CoverImage: book.CoverImage,
		CreatedAt:  book.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  book.UpdatedAt.Format("2006-01-02 15:04:05"),
		Stock: StockResponse{
			TotalStock:     totalStock,
			AvailableStock: availableStock,
		},
	}
}

func ToBookResponseList(books []entity.Book) []BookResponse {
	var responseList []BookResponse
	for _, book := range books {
		responseList = append(responseList, BookResponse{
			ID:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Publisher: book.Publisher,
			Category:  book.Category,
		})
	}
	return responseList
}
