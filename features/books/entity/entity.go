package entity

type Book struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	Category   string `json:"category"`
	Isbn       string `json:"isbn"`
	Year       string `json:"year"`
	CoverImage string `json:"cover_image"`
}