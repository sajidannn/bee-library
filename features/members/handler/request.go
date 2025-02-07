package handler

type MemberCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Photo    string `json:"photo"`
}

type MemberUpdateRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
	Photo    *string `json:"photo"`
}
