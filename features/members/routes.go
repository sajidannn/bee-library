package members

import (
	"bee-library/db"
	"bee-library/features/members/handler"
	"bee-library/features/members/repository"
	"bee-library/features/members/service"

	"github.com/gin-gonic/gin"
)

func RegisterMemberRoutes(r *gin.Engine) {
	memberRepo := repository.NewMemberRepository(db.DB)
	memberService := service.NewMemberService(memberRepo)
	memberHandler := handler.NewMemberHandler(memberService)

	r.GET("/members", memberHandler.GetAllMembers)
	r.GET("/members/:id", memberHandler.GetMemberByID)
	r.POST("/members", memberHandler.CreateMember)
	r.PUT("/members/:id", memberHandler.UpdateMember)
	r.DELETE("/members/:id", memberHandler.DeleteMember)
}
