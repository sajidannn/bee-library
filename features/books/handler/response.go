package handler

import "bee-library/features/books/entity"

type BookResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	Category   string `json:"category"`
	Isbn       string `json:"isbn"`
	Year       string `json:"year"`
	CoverImage string `json:"cover_image,omitempty"`
}

func ToBookResponse(book entity.Book) BookResponse {
	return BookResponse{
		ID:         book.ID,
		Title:      book.Title,
		Author:     book.Author,
		Publisher:  book.Publisher,
		Category:   book.Category,
		Isbn:       book.Isbn,
		Year:       book.Year,
		CoverImage: book.CoverImage,
	}
}

func ToBookResponseList(books []entity.Book) []BookResponse {
	var responseList []BookResponse
	for _, book := range books {
		responseList = append(responseList, ToBookResponse(book))
	}
	return responseList
}