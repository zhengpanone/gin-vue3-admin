package system

import (
	v1 "gin-api-learn/api/v1"
	"gin-api-learn/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(engine *gin.Engine) {
	userRouter := engine.Group("v1/user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

	{
		userRouter.POST("register", baseApi.Register)             // 用户注册
		userRouter.POST("changePassword", baseApi.ChangePassword) //用户修改密码
		userRouter.POST("login", baseApi.Login)                   // 用户登录

	}
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", baseApi.GetUser)
	}

}
