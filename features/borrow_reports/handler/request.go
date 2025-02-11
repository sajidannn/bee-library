package handler

type GetBorrowReportsRequest struct {
	BookID    *uint   `form:"book_id"`
	StartDate *string `form:"start_date"`
	EndDate   *string `form:"end_date"`
}
