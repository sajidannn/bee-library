package handler

type GetReturnReportsRequest struct {
	BookID    *uint   `form:"book_id"`
	MemberID  *uint   `form:"member_id"`
	StartDate *string `form:"start_date"`
	EndDate   *string `form:"end_date"`
}