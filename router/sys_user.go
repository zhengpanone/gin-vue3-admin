package router

import (
	"gin-api-learn/middleware"

	"gin-api-learn/api/v1/system"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(engine *gin.Engine) {
	userRouter := engine.Group("v1/user")

	{

		userRouter.POST("register", system.Register)             // 用户注册
		userRouter.POST("changePassword", system.ChangePassword) //用户修改密码
		userRouter.POST("login", system.Login)                   // 用户登录

	}
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", system.GetUser)
	}

}
