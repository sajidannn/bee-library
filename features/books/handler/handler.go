package handler

import (
	"bee-library/features/books/entity"
	"bee-library/helper"
	"bee-library/utils"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service entity.BookService
}

func NewBookHandler(service entity.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError{
			Status:  "error",
			Message: "Failed to fetch books",
		})
		return
	}
	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "There's no data"})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Books retrieved successfully",
		Data:    ToBookResponseList(books),
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book,totalStock, availableStock, err := h.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Book retrieved successfully",
		Data:    ToBookDetailResponse(*book, totalStock, availableStock),
	})
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req BookCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	var coverURL string
	file, fileExists := c.Get("cover_image_file")
	fileName, nameExists := c.Get("cover_image_fileName")

	if fileExists && nameExists {
		uploadedURL, err := utils.UploadToCloudinary(file.(multipart.File), fileName.(string), "book-covers")
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseError{
				Status:  "error",
				Message: "Failed to upload photo",
			})
			return
		}
		coverURL = uploadedURL
	}

	newBook := entity.Book{
		Title:      req.Title,
		Author: 	 	req.Author,
		Publisher:  req.Publisher,
		Category:   req.Category,
		Isbn:       req.Isbn,
		Year:       req.Year,
		CoverImage: coverURL,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := h.service.CreateBook(&newBook)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, helper.Response{
		Status:  "success",
		Message: "Book created successfully",
		Data:    ToBookResponse(newBook),
	})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req BookUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	existingBook, _, _, err := h.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	updatedBook := entity.Book{
		UpdatedAt: time.Now(),
	}
	if req.Title != nil {
		updatedBook.Title = *req.Title
	}
	if req.Author != nil {
		updatedBook.Author = *req.Author
	}
	if req.Publisher != nil {
		updatedBook.Publisher = *req.Publisher
	}
	if req.Category != nil {
		updatedBook.Category = *req.Category
	}
	if req.Year != nil {
		updatedBook.Year = *req.Year
	}

	file, fileExists := c.Get("cover_image_file")
	fileName, nameExists := c.Get("cover_image_fileName")

	if fileExists && nameExists {
		if existingBook.CoverImage != "" {
			err := utils.DeleteFromCloudinary(existingBook.CoverImage)
			if err != nil {
				c.JSON(http.StatusInternalServerError, helper.ResponseError{
					Status:  "error",
					Message: "Failed to delete old photo",
				})
				return
			}
		}
		uploadedURL, err := utils.UploadToCloudinary(file.(multipart.File), fileName.(string), "book-covers")
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseError{
				Status:  "error",
				Message: "Failed to upload photo",
			})
			return
		}
		updatedBook.CoverImage = uploadedURL
	}

	err = h.service.UpdateBook(uint(id), &updatedBook)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Book updated successfully",
	})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteBook(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Book deleted successfully",
	})
}