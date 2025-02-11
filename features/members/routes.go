package members

import (
	"bee-library/db"
	"bee-library/features/members/handler"
	"bee-library/features/members/repository"
	"bee-library/features/members/service"
	middleware "bee-library/middlewares"

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
		memberRoutes.POST("", middleware.FileUploadMiddleware("photo"), memberHandler.CreateMember)
		memberRoutes.PUT("/:id", middleware.FileUploadMiddleware("photo"), memberHandler.UpdateMember)
		memberRoutes.DELETE("/:id", memberHandler.DeleteMember)
	}
}
