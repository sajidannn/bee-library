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
	Isbn       string        `json:"isbn,omitempty"`
	Year       string        `json:"year,omitempty"`
	CoverImage string        `json:"cover_image,omitempty"`
	CreatedAt  string        `json:"created_at,omitempty"`
	UpdatedAt  string        `json:"updated_at,omitempty"`
	Stock      StockResponse `json:"stock,omitempty"`
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
	var createdAtStr, updatedAtStr string
	if !book.CreatedAt.IsZero() {
		createdAtStr = book.CreatedAt.Format("2006-01-02 15:04:05")
	}
	if !book.UpdatedAt.IsZero() {
		updatedAtStr = book.UpdatedAt.Format("2006-01-02 15:04:05")
	}

	return BookDetailResponse{
		ID:         book.ID,
		Title:      book.Title,
		Author:     book.Author,
		Publisher:  book.Publisher,
		Category:   book.Category,
		Isbn:       book.Isbn,
		Year:       book.Year,
		CoverImage: book.CoverImage,
		CreatedAt:  createdAtStr,
		UpdatedAt:  updatedAtStr,
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
