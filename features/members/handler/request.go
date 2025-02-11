package handler

type MemberCreateRequest struct {
	Name    string `form:"name" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Phone   string `form:"phone" binding:"required"`
	Address string `form:"address" binding:"required"`
}

type MemberUpdateRequest struct {
	Name    *string `form:"name"`
	Phone   *string `form:"phone"`
	Address *string `form:"address"`
}
