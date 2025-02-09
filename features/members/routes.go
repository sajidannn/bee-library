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

	memberRoutes := r.Group("/members")
	{
		memberRoutes.GET("/", memberHandler.GetAllMembers)
		memberRoutes.GET("/:id", memberHandler.GetMemberByID)
		memberRoutes.POST("/", memberHandler.CreateMember)
		memberRoutes.PUT("/:id", memberHandler.UpdateMember)
		memberRoutes.DELETE("/:id", memberHandler.DeleteMember)
	}
}
