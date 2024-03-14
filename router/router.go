package router

import (
	"blockaction_homework/app/controller"
	"blockaction_homework/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 初始化所有的路徑
func Init(server *gin.Engine, db *gorm.DB) {
	routerGroup := server.Group("/api")
	InitMemberController(routerGroup, db)
}

// 初始化會員控制器
func InitMemberController(r *gin.RouterGroup, db *gorm.DB) {
	memberService := service.NewMemberService(db)
	memberController := controller.NewMemberController(memberService)
	r.POST("member", memberController.Register)
	r.POST("session", memberController.Login)
	r.GET("member/:username", memberController.GetMemberData)
	r.GET("members", memberController.GetAllMemberData)
}