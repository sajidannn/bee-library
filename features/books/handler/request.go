package handler

type BookCreateRequest struct {
	Title     string `form:"title" binding:"required"`
	Author    string `form:"author" binding:"required"`
	Publisher string `form:"publisher"`
	Category  string `form:"category"`
	Isbn      string `form:"isbn"`
	Year      string `form:"year"`
	// CoverImage string `json:"cover_image"`
}

type BookUpdateRequest struct {
	Title     *string `form:"title"`
	Author    *string `form:"author"`
	Publisher *string `form:"publisher"`
	Category  *string `form:"category"`
	Year      *string `form:"year"`
	// CoverImage *string `json:"cover_image"`
}