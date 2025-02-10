package handler

type BookCreateRequest struct {
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	Publisher  string `json:"publisher"`
	Category   string `json:"category"`
	Isbn       string `json:"isbn"`
	Year       string `json:"year"`
	CoverImage string `json:"cover_image"`
}

type BookUpdateRequest struct {
	Title      *string `json:"title"`
	Author     *string `json:"author"`
	Publisher  *string `json:"publisher"`
	Category   *string `json:"category"`
	Year       *string `json:"year"`
	CoverImage *string `json:"cover_image"`
}